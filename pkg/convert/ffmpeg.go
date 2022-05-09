package convert

import (
	"errors"
	"os"

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
)

func (codec FFMPEGCodec) ContentType() string {

	const imageCodec FFMPEGCodec = "image/"

	var contentType FFMPEGCodec

	switch codec {

	case FFMPEGCodecWebp:
		contentType = imageCodec + FFMPEGCodecWebp

	case FFMPEGCodecPng:
		contentType = imageCodec + FFMPEGCodecPng

	default:
		return ""
	}

	return string(contentType)
}

func (codec FFMPEGCodec) FileExtension() string {

	var contentType FFMPEGCodec

	switch codec {

	case FFMPEGCodecWebp:
		contentType = FFMPEGCodecWebp

	case FFMPEGCodecPng:
		contentType = FFMPEGCodecPng

	default:
		return ""
	}

	return string(contentType)
}

type FFMPEGScale struct {
	Width  int
	Height int
}

const (
	FFMPEGScaleAuto = -1
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
