package convert

import "strings"

type Transformation struct {
	Format     FFMPEGFormat
	VideoCodec FFMPEGCodec
	AudioCodec FFMPEGCodec
	Scale      *FFMPEGScale

	outputCacheFile string
}

func NewTransformation() *Transformation {
	return &Transformation{
		Format:     FFMPEGFormatNone,
		VideoCodec: FFMPEGCodecNone,
		AudioCodec: FFMPEGCodecNone,
		Scale:      nil,
	}
}

func (tr *Transformation) ScaleByWidth(width int) *Transformation {

	if tr.Scale == nil {
		tr.Scale = &FFMPEGScale{
			Width:  FFMPEGScaleAuto,
			Height: FFMPEGScaleAuto,
		}
	}

	tr.Scale.Width = width
	return tr
}

func (tr *Transformation) ScaleByHeight(height int) *Transformation {

	if tr.Scale == nil {
		tr.Scale = &FFMPEGScale{
			Width:  FFMPEGScaleAuto,
			Height: FFMPEGScaleAuto,
		}
	}

	tr.Scale.Height = height
	return tr
}

func (tr *Transformation) NoScaling() *Transformation {
	tr.Scale = nil
	return tr
}

// Detect the processing format from the content type. This is only used to detect input file format.
// Detected format wont work for output file formats.
//
// Returns FFMPEGFormat in pipe mode. Use FFMPEGFormat.ToFile() to get the corresponding file mode.
func FormatFromContentType(contentType string) FFMPEGFormat {

	contentType = strings.TrimSpace(strings.ToLower(contentType))

	switch contentType {
	case
		"image/webp",
		"image/png",
		"image/jpeg", "image/jpg",
		"image/gif":
		return FFMPEGFormatImagePipe

	default:
		return FFMPEGFormatNone
	}

}

// Set the content type for transformation
func (tr *Transformation) ContentType(contentType string) *Transformation {
	tr.Format = FormatFromContentType(contentType)

	switch contentType {
	case "image/webp":
		tr.VideoCodec = FFMPEGCodecWebp
		tr.AudioCodec = FFMPEGCodecNone

	case "image/png":
		tr.VideoCodec = FFMPEGCodecPng
		tr.AudioCodec = FFMPEGCodecNone

	case "image/jpeg", "image/jpg":
		tr.VideoCodec = FFMPEGCodecJpg
		tr.AudioCodec = FFMPEGCodecNone

	default:
		tr.VideoCodec = FFMPEGCodecNone
		tr.AudioCodec = FFMPEGCodecNone

	}

	return tr
}

func (tr *Transformation) GetFileExtention() string {

	if tr.VideoCodec != FFMPEGCodecNone {
		return tr.VideoCodec.FileExtension()

	} else if tr.AudioCodec != FFMPEGCodecNone {
		return tr.AudioCodec.FileExtension()

	} else {
		return ""

	}
}
