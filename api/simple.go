package api

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/HasinduLanka/diviyago/pkg/convert"
)

func SimpleEndpoint(wr http.ResponseWriter, req *http.Request) {

	fileBytes, fileBytesErr := os.ReadFile(`TestMedia/go.png`)

	if fileBytesErr != nil {
		log.Panicln(`/api/webp : file read error : `, fileBytesErr)
		wr.Write([]byte(`/api/webp : file read error : ` + fileBytesErr.Error()))
		return
	}

	converter := convert.NewImageConverter()
	img_s := converter.AddTransformation(convert.NewTransformation().ContentType(`image/webp`).ScaleByWidth(128))
	img_m := converter.AddTransformation(convert.NewTransformation().ContentType(`image/webp`).ScaleByWidth(640))
	img_l := converter.AddTransformation(convert.NewTransformation().ContentType(`image/webp`).ScaleByWidth(1080))
	img_h := converter.AddTransformation(convert.NewTransformation().ContentType(`image/webp`).ScaleByWidth(1920))

	result := converter.Convert(fileBytes, nil)

	if result.Error != nil {
		log.Panicln(`/api/webp : convert error : `, result.Error)
		wr.Write([]byte(`/api/webp : convert error : ` + result.Error.Error()))
		return
	}

	transformedIDs := []string{img_s, img_m, img_l, img_h}

	// show a transformed image according to the current time second
	indx := time.Now().Second() % len(transformedIDs)

	selectedRes := result.TransformedResults[transformedIDs[indx]]

	wr.Header().Set(`Content-Type`, selectedRes.VideoCodec.ContentType())
	wr.Write(selectedRes.Data)
}
