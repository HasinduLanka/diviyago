package convert

import (
	"errors"
	"os"
	"strconv"

	"github.com/HasinduLanka/diviyago/pkg/goex"
)

type FFMPEGFormat string

const (
	FFMPEGFormatNone      FFMPEGFormat = ""
	FFMPEGFormatImagePipe FFMPEGFormat = "image2pipe"
	FFMPEGFormatImageFile FFMPEGFormat = "image2"
)

func (fm FFMPEGFormat) ToPipe() FFMPEGFormat {

	switch fm {

	case FFMPEGFormatImageFile:
		return FFMPEGFormatImagePipe

	default:
		return fm
	}
}

func (fm FFMPEGFormat) ToFile() FFMPEGFormat {

	switch fm {

	case FFMPEGFormatImagePipe:
		return FFMPEGFormatImageFile

	default:
		return fm
	}
}

type FFMPEGCodec string

const (
	FFMPEGCodecNone FFMPEGCodec = ""
	FFMPEGCodecWebp FFMPEGCodec = "webp"
	FFMPEGCodecPng  FFMPEGCodec = "png"
	FFMPEGCodecJpg  FFMPEGCodec = "mjpeg"
)

type FFMPEGCodecType string

const (
	FFMPEGCodecTypeVideo FFMPEGCodecType = "v"
	FFMPEGCodecTypeAudio FFMPEGCodecType = "a"
)

func (codec FFMPEGCodec) ContentType() string {

	const imageCodec = "image/"

	switch codec {

	case FFMPEGCodecWebp, FFMPEGCodecPng, FFMPEGCodecJpg:
		return imageCodec + codec.FileExtension()

	default:
		return ""
	}

}

func (codec FFMPEGCodec) FileExtension() string {

	var contentType FFMPEGCodec

	switch codec {

	case FFMPEGCodecWebp:
		contentType = "webp"

	case FFMPEGCodecPng:
		contentType = "png"

	case FFMPEGCodecJpg:
		contentType = "jpg"

	default:
		return ""
	}

	return string(contentType)
}

type TransformResolution struct {
	Width  int
	Height int
}

const (
	TransformResolutionAuto = -1
)

// Extracts FFMPEG executable and dependencies from cache.
// Returns FFMPEG executable path
func InitializeFFMPEG() (string, error) {

	tmpDir := os.TempDir() + `/`

	cacheDir := tmpDir + `diviyago/`

	extractErr := goex.DeployEmbedFiles(cacheDir)

	if extractErr != nil {
		return "", errors.New("Extracting FFMPEG failed : " + extractErr.Error())
	}

	ffmpegPath := cacheDir + `EmbedFiles/ffmpeg-linux-amd64/ffmpeg`

	return ffmpegPath, nil

}

// Image / Audio / Video quality parameter for FFMPEG.
// Must be in range of 0 to 100, where 0 is lowest and 100 is highest quality.
//
// FFMPEGQualityHigh (80) is the default value.
type TransformQuality int

const (
	TransformQualityNone  TransformQuality = 0
	TransformQualityLow   TransformQuality = 55
	TransformQualityMid   TransformQuality = 70
	TransformQualityHigh  TransformQuality = 80
	TransformQualityExtra TransformQuality = 90
)

// Returns the quality parameter for FFMPEG in 0 to 100 scale. Useful for formats like WEBP
func (ql TransformQuality) To100Scale() string {
	return strconv.Itoa(int(ql))
}

// Returns the quality parameter for FFMPEG Q-Scale. Value ranges from 0 to 31, where 31 is the lowest quality.
func (ql TransformQuality) ToQScale() string {
	return strconv.Itoa(int((100 - ql) * 31 / 100))
}

// Returns the quality, compression selection command line arguments for FFMPEG, depending on the FFMPEG Codec.
//
// Examples:
//
// for MP4, the return value could be [ -qscale:v 3 ]
//
// for WEBP, the return value could be [ -quality 80 -compression_level 6 ]
func (ql TransformQuality) ToArgs(codec FFMPEGCodec, codecType FFMPEGCodecType) []string {

	switch codec {
	case FFMPEGCodecWebp:
		// Asumption: it is always a video codec
		return []string{"-quality", ql.To100Scale()}

	default:
		return []string{"-qscale:" + string(codecType), ql.ToQScale()}
	}

}
