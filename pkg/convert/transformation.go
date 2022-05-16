package convert

import "strings"

type Transformation struct {
	Format     FFMPEGFormat
	VideoCodec FFMPEGCodec
	AudioCodec FFMPEGCodec
	Resolution *TransformResolution
	Quality    TransformQuality

	outputCacheFile string
}

func NewTransformation() *Transformation {
	return &Transformation{
		Format:     FFMPEGFormatNone,
		VideoCodec: FFMPEGCodecNone,
		AudioCodec: FFMPEGCodecNone,
		Resolution: nil,
		Quality:    TransformQualityHigh,
	}
}

func (tr *Transformation) ResolutionByWidth(width int) *Transformation {

	if tr.Resolution == nil {
		tr.Resolution = &TransformResolution{
			Width:  TransformResolutionAuto,
			Height: TransformResolutionAuto,
		}
	}

	tr.Resolution.Width = width
	return tr
}

func (tr *Transformation) ResolutionByHeight(height int) *Transformation {

	if tr.Resolution == nil {
		tr.Resolution = &TransformResolution{
			Width:  TransformResolutionAuto,
			Height: TransformResolutionAuto,
		}
	}

	tr.Resolution.Height = height
	return tr
}

func (tr *Transformation) NoScaling() *Transformation {
	tr.Resolution = nil
	return tr
}

func (tr *Transformation) SetQuality(quality TransformQuality) *Transformation {
	tr.Quality = quality
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

		if strings.HasPrefix(contentType, "image") {
			tr.VideoCodec = FFMPEGCodecWebp
			tr.AudioCodec = FFMPEGCodecNone

		} else {
			tr.VideoCodec = FFMPEGCodecNone
			tr.AudioCodec = FFMPEGCodecNone

		}
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

func (tr *Transformation) GetContentType() string {

	if tr.VideoCodec != FFMPEGCodecNone {
		return tr.VideoCodec.ContentType()

	} else if tr.AudioCodec != FFMPEGCodecNone {
		return tr.AudioCodec.ContentType()

	} else {
		return ""

	}
}
