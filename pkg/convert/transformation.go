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

func FormatFromContentType(contentType string) FFMPEGFormat {

	if strings.HasPrefix(contentType, "image/") {
		return FFMPEGFormatImagePipe

	} else {
		return FFMPEGFormatNone

	}
}

func (tr *Transformation) ContentType(contentType string) *Transformation {
	tr.Format = FormatFromContentType(contentType)

	switch contentType {
	case "image/webp":
		tr.VideoCodec = FFMPEGCodecWebp
		tr.AudioCodec = FFMPEGCodecNone

	case "image/png":
		tr.VideoCodec = FFMPEGCodecPng
		tr.AudioCodec = FFMPEGCodecNone

	default:
		tr.VideoCodec = FFMPEGCodecNone
		tr.AudioCodec = FFMPEGCodecNone

	}

	return tr
}

// (format FFMPEGFormat, videoCodec FFMPEGCodec, audioCodec FFMPEGCodec)
