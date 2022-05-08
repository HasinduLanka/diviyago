package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HasinduLanka/diviyago/goex"
)

func SimpleEndpoint(wr http.ResponseWriter, req *http.Request) {

	result, exeErr := goex.ExcecProgramToString("diviya")

	if exeErr != nil {
		log.Panicln("/api/simple : exe error : ", exeErr)
		wr.Write([]byte("/api/simple : exe error : " + exeErr.Error()))
		return
	}

	wr.Write([]byte("Simple Endpoint " + fmt.Sprint(result)))
}
