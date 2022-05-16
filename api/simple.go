package api

import (
	"log"
	"net/http"
	"time"

	"github.com/HasinduLanka/diviyago/pkg/convert"
	"github.com/HasinduLanka/diviyago/testmedia"
)

func SimpleEndpoint(wr http.ResponseWriter, req *http.Request) {

	testFileBytes := testmedia.FileGoLogo

	converter := convert.NewImageConverter()
	img_s := converter.AddTransformation(convert.NewTransformation().ContentType(`image/webp`).ResolutionByWidth(128))
	img_m := converter.AddTransformation(convert.NewTransformation().ContentType(`image/webp`).ResolutionByWidth(640))
	img_l := converter.AddTransformation(convert.NewTransformation().ContentType(`image/webp`).ResolutionByWidth(1080))
	img_h := converter.AddTransformation(convert.NewTransformation().ContentType(`image/webp`).ResolutionByWidth(1920))

	result := converter.Convert(testFileBytes, nil)

	if result.Error != nil {
		log.Panicln(`/api/simple : convert error : `, result.Error)
		wr.Write([]byte(`/api/simple : convert error : ` + result.Error.Error()))
		return
	}

	transformedIDs := []string{img_s, img_m, img_l, img_h}

	// show a transformed image according to the current time second
	indx := time.Now().Second() % len(transformedIDs)

	selectedRes := result.TransformedResults[transformedIDs[indx]]

	wr.Header().Set(`Content-Type`, selectedRes.VideoCodec.ContentType())
	wr.Write(*selectedRes.Data)
}
