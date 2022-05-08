package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/HasinduLanka/diviyago/goex"
)

func SimpleEndpoint(wr http.ResponseWriter, req *http.Request) {

	libwasmerBytes, libwasmerErr := goex.GetFile("libwasmer.so")

	if libwasmerErr != nil {
		log.Fatal(libwasmerErr)
	}

	os.WriteFile("libwasmer.so", libwasmerBytes, 0644)

	os.Mkdir("packaged/lib/linux-amd64/", 0777)

	os.WriteFile("packaged/lib/linux-amd64/libwasmer.so", libwasmerBytes, 0644)

	result, exeErr := goex.ExcecProgramToString("diviya")

	if exeErr != nil {
		log.Panicln("/api/simple : exe error : ", exeErr)
		wr.Write([]byte("/api/simple : exe error : " + exeErr.Error()))
		return
	}

	wr.Write([]byte("Simple Endpoint " + fmt.Sprint(result)))
}
