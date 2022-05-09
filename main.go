package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/HasinduLanka/diviyago/api"
	"github.com/HasinduLanka/diviyago/pkg/wasm"
	wasmer "github.com/wasmerio/wasmer-go/wasmer"
)

// This source file faciliates a simple web server that serves static files and handles requests to the /api/hello endpoint.
// This will not be executed on Vercel.

func main() {

	multiplexer := &http.ServeMux{}

	// serve static files
	multiplexer.Handle("/", http.FileServer(http.Dir("./frontend/public")))

	multiplexer.HandleFunc("/api/hello", api.HelloEndpoint)
	multiplexer.HandleFunc("/api/simple", api.SimpleEndpoint)
	multiplexer.HandleFunc("/api/webp", api.WebpEndpoint)

	log.Println("Listening on port 31603. Visit http://localhost:31603 if you're running this locally.")

	go local()

	// Blocks until the program is terminated
	serveErr := http.ListenAndServe(":31603", multiplexer)

	// serveErr is always non nil
	log.Println(serveErr.Error())

}

func simplewasm() {
	wasmBytes, wasmFileReadErr := ioutil.ReadFile("./wasmfiles/simple.wasm")

	if wasmFileReadErr != nil {
		log.Fatal(wasmFileReadErr)
	}

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	// Compiles the module
	module, _ := wasmer.NewModule(store, wasmBytes)

	// Instantiates the module
	importObject := wasmer.NewImportObject()
	instance, _ := wasmer.NewInstance(module, importObject)

	// Gets the `sum` exported function from the WebAssembly instance.
	sum, _ := instance.Exports.GetFunction("sum")

	// Calls that exported function with Go standard values. The WebAssembly
	// types are inferred and values are casted automatically.
	result, _ := sum(5, 37)

	fmt.Println(result) // 42!
}

func local() {
	instance, instanceErr := wasm.GetWASM("simple.wasm")

	if instanceErr != nil {
		log.Fatal(instanceErr)
	}

	// Gets the `sum` exported function from the WebAssembly instance.
	sum, _ := instance.Exports.GetFunction("sum")

	// Calls that exported function with Go standard values. The WebAssembly
	// types are inferred and values are casted automatically.
	result, _ := sum(50, 307)

	fmt.Println(result)

}
