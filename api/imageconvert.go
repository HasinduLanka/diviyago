package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/HasinduLanka/diviyago/pkg/convert"
)

func ImageConvertEndpoint(wr http.ResponseWriter, req *http.Request) {

	// Read the request body as a byte array

	uploadReqData, errUploadReqData := ioutil.ReadAll(req.Body)

	if errUploadReqData != nil {
		writeError(wr, "/api/upload : request body is not readable : "+errUploadReqData.Error())
		return
	}

	if len(uploadReqData) == 0 {
		writeError(wr, "/api/upload : request body is empty. Send a file to upload")
		return
	}

	contentType := req.Header.Get(`Content-Type`)

	if contentType == `` {
		writeError(wr, "/api/upload : request header is missing content-type")
		return
	}

	converter, converterCtorErr := convert.FromContentType(contentType)

	if converterCtorErr != nil {
		writeError(wr, "/api/upload : converter constructor error : "+converterCtorErr.Error())
		return
	}

	converter.AddTransformation(convert.NewTransformation().ContentType(`image/webp`).ScaleByWidth(128))
	converter.AddTransformation(convert.NewTransformation().ContentType(`image/webp`).ScaleByWidth(640))
	converter.AddTransformation(convert.NewTransformation().ContentType(`image/webp`).ScaleByWidth(1080))
	converter.AddTransformation(convert.NewTransformation().ContentType(`image/webp`).ScaleByWidth(1920))

	result := converter.Convert(uploadReqData, nil)

	if result.Error != nil {
		writeError(wr, `/api/upload : convert error : `+result.Error.Error())
		return
	}

	resJson, resJsonErr := json.MarshalIndent(result, "", "  ")

	if resJsonErr != nil {
		writeError(wr, `/api/upload : json encode error : `+resJsonErr.Error())
		return
	}

	wr.Header().Set(`Content-Type`, `application/json`)
	wr.Write(resJson)

}

func writeError(wr http.ResponseWriter, errMsg string) {
	log.Println(errMsg)

	jsonErrMsg, _ := json.MarshalIndent(map[string]string{`error`: errMsg}, "", "  ")

	http.Error(wr, string(jsonErrMsg), http.StatusBadRequest)

}
