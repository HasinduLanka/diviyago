package wasm

import (
	"io/ioutil"
	"log"

	wasmer "github.com/wasmerio/wasmer-go/wasmer"
)

var wasmBytesCache = make(map[string]*[]byte)

func GetWASM(wasmFileName string) (*wasmer.Instance, error) {
	if wasmBytes, ok := wasmBytesCache[wasmFileName]; ok {
		return getWASMFromBytes(wasmBytes)
	}

	wasmBytes, err := ioutil.ReadFile(wasmFileName)
	if err != nil {
		return nil, err
	}

	wasmBytesCache[wasmFileName] = &wasmBytes

	return getWASMFromBytes(&wasmBytes)
}

func getWASMFromBytes(wasmBytes *[]byte) (*wasmer.Instance, error) {

	module, moduleCompileErr := getWASMModuleFromBytes(wasmBytes)

	if moduleCompileErr != nil {
		return nil, moduleCompileErr
	}

	// Instantiates the module
	importObject := wasmer.NewImportObject()
	instance, instanciateErr := wasmer.NewInstance(module, importObject)

	if instanciateErr != nil {
		log.Println("WASM instantiate err: ", instanciateErr.Error())
		return nil, instanciateErr
	}

	return instance, nil
}

func getWASMModuleFromBytes(wasmBytes *[]byte) (*wasmer.Module, error) {
	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	// Compiles the module
	module, moduleCompileErr := wasmer.NewModule(store, *wasmBytes)

	if moduleCompileErr != nil {
		log.Println("WASM module compile err: ", moduleCompileErr.Error())
		return nil, moduleCompileErr
	}

	return module, nil
}
