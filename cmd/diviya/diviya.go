package main

import (
	"fmt"
	"log"

	"github.com/HasinduLanka/diviyago/pkg/wasm"
)

func main() {
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
