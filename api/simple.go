package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HasinduLanka/diviyago/pkg/wasm"
)

func SimpleEndpoint(wr http.ResponseWriter, req *http.Request) {

	instance, instanceErr := wasm.GetWASM("simple.wasm")

	if instanceErr != nil {
		log.Println("/api/simple : wasm instanciate error : " + instanceErr.Error())
		wr.Write([]byte("/api/simple : wasm instanciate error : " + instanceErr.Error()))

	}

	// Gets the `sum` exported function from the WebAssembly instance.
	sum, getFuncErr := instance.Exports.GetFunction("sum")

	if getFuncErr != nil {
		log.Println("/api/simple : wasm get func error : " + getFuncErr.Error())
		wr.Write([]byte("/api/simple : wasm get func error : " + getFuncErr.Error()))
		return
	}

	// Calls that exported function with Go standard values. The WebAssembly
	// types are inferred and values are casted automatically.
	result, resultErr := sum(50, 307)

	if resultErr != nil {
		log.Println("/api/simple : wasm func error : " + resultErr.Error())
		wr.Write([]byte("/api/simple : wasm func error : " + resultErr.Error()))
		return
	}

	log.Println(result)

	wr.Write([]byte("Simple Endpoint " + fmt.Sprint(result)))
}
