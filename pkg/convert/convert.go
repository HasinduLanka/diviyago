package convert

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/HasinduLanka/diviyago/pkg/goex"
	"github.com/google/uuid"
)

const ffmpegStdIOPipe = "pipe:"

type Converter struct {
	InputFormat FFMPEGFormat

	// Transformation to be applied. map[ uniqueTransformationID ] *Transformation
	Transform map[string]*Transformation
}

type TransformationResult struct {
	Transformation
	Data []byte
}

type ConvertionResult struct {
	Success bool
	Error   error

	TransformedResults map[string]*TransformationResult
}

func NewConverter() *Converter {

	return &Converter{
		Transform: make(map[string]*Transformation),
	}
}

func NewImageConverter() *Converter {

	return &Converter{
		InputFormat: FFMPEGFormatImagePipe,
		Transform:   make(map[string]*Transformation),
	}
}

// Create a converter according to a content type.
func FromContentType(contentType string) (*Converter, error) {

	// Test for content type
	fm := FormatFromContentType(contentType)

	switch fm {

	case FFMPEGFormatImagePipe:
		return NewImageConverter(), nil

	case FFMPEGFormatImageFile:
		return NewConverter(), nil

	default:
		return nil, errors.New("Unknown content type: " + contentType)

	}

}

// Add transformation to the converter. Returns a unique transformation ID.
func (cnv *Converter) AddTransformation(t *Transformation) string {
	uniqueID := uuid.New().String()
	cnv.Transform[uniqueID] = t
	return uniqueID
}

func (cnv *Converter) Convert(input []byte, endtask chan bool) ConvertionResult {

	ffmpegExe, initErr := InitializeFFMPEG()

	if initErr != nil {
		return ConvertionResult{Error: initErr}
	}

	// generate UUID for output directory
	outputDir := os.TempDir() + "/" + uuid.New().String() + "/"

	// create output directory
	os.MkdirAll(outputDir, os.ModePerm)

	defer func() {
		// delete output directory
		os.RemoveAll(outputDir)
	}()

	cmd, cmdargs := cnv.buildCommand(outputDir, ffmpegExe)

	log.Println("FFMPEG command:", cmd, cmdargs)

	// run command
	cmdRes, cmdRunErr := goex.ExcecTask(endtask, input, cmd, cmdargs...)

	if cmdRunErr != nil {
		return ConvertionResult{Error: cmdRunErr}
	}

	result := ConvertionResult{
		Success:            true,
		TransformedResults: make(map[string]*TransformationResult),
	}

	for tid, trn := range cnv.Transform {
		if trn.outputCacheFile == ffmpegStdIOPipe {

			result.TransformedResults[tid] = &TransformationResult{
				Transformation: *trn,
				Data:           cmdRes,
			}

		} else {

			// Read output fileBytes
			fileBytes, fileReadErr := os.ReadFile(trn.outputCacheFile)

			if fileReadErr != nil {
				result.Error = errors.New("Failed to read output file: " + tid + ": " + fileReadErr.Error())
				result.Success = false
				return result
			}

			result.TransformedResults[tid] = &TransformationResult{
				Transformation: *trn,
				Data:           fileBytes,
			}
		}
	}

	return result
}

func (cnv *Converter) buildCommand(outputDir string, ffmpegExe string) (string, []string) {

	command := make([]string, 0, 32)

	command = append(command, "-y")

	if len(cnv.InputFormat) > 0 {
		command = append(command, "-f", string(cnv.InputFormat))
	}

	command = append(command, "-i", ffmpegStdIOPipe)

	outputPipeAvailable := true

	for trnsID, trn := range cnv.Transform {

		if trn.Scale != nil {
			command = append(command, `-vf`, `scale=`+strconv.Itoa(trn.Scale.Width)+`:`+strconv.Itoa(trn.Scale.Height))
		}

		if trn.Format != FFMPEGFormatNone {
			command = append(command, `-f`, string(trn.Format))
		}

		if trn.VideoCodec != FFMPEGCodecNone {
			command = append(command, `-c:v`, string(trn.VideoCodec))
		}

		if trn.AudioCodec != FFMPEGCodecNone {
			command = append(command, `-c:a`, string(trn.AudioCodec))
		}

		if len(trn.outputCacheFile) == 0 {

			if outputPipeAvailable {
				trn.Format = trn.Format.ToPipe()
				trn.outputCacheFile = ffmpegStdIOPipe
				outputPipeAvailable = false

			} else if trn.VideoCodec != FFMPEGCodecNone {
				trn.Format = trn.Format.ToFile()
				trn.outputCacheFile = outputDir + trnsID + "." + trn.VideoCodec.FileExtension()

			} else if trn.AudioCodec != FFMPEGCodecNone {
				trn.Format = trn.Format.ToFile()
				trn.outputCacheFile = outputDir + trnsID + "." + trn.AudioCodec.FileExtension()

			} else {
				trn.Format = trn.Format.ToFile()
				trn.outputCacheFile = outputDir + trnsID

			}
		}

		command = append(command, trn.outputCacheFile)
	}

	return ffmpegExe, command

}
