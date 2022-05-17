package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/HasinduLanka/diviyago/pkg/convert"
	"github.com/HasinduLanka/diviyago/testmedia"
	"github.com/aws/aws-lambda-go/lambda"
)

// Build with
// GOOS=linux CGO_ENABLED=0 go build  -o ./lambda/build/ ./lambda/simple
//
// ZIP the binary file './lambda/build/simple'
// Upload it to AWS Lambda

type SimpleResponse struct {
	Data        []byte `json:"data"`
	ContentType string `json:"contentType"`
}

func Handler(ctx context.Context) (string, error) {

	testFileBytes := testmedia.FileEarth

	converter := convert.NewImageConverter()
	img_m := converter.AddTransformation(convert.NewTransformation().ContentType(`image/webp`).ResolutionByWidth(640))

	result := converter.Convert(testFileBytes, nil)

	if result.Error != nil {
		errLine := `/api/simple : convert error : ` + result.Error.Error()
		log.Println(errLine)
		return errLine, result.Error
	}

	selectedRes := result.TransformedResults[img_m]

	resp := SimpleResponse{
		Data:        *selectedRes.Data,
		ContentType: selectedRes.GetContentType(),
	}

	jsonResp, jsonErr := json.Marshal(resp)

	if jsonErr != nil {
		errLine := `/api/simple : json marshal error : ` + jsonErr.Error()
		log.Println(errLine)
		return errLine, jsonErr
	}

	return string(jsonResp), nil
}

func main() {
	lambda.Start(Handler)
}
