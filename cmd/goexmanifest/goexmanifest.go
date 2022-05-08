package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/HasinduLanka/diviyago/pkg/symembed"
)

func main() {
	log.Println("goexmanifest")

	manifest, manifestErr := symembed.GenManifest("goex/exeFiles")

	if manifestErr != nil {
		log.Panicln("goexmanifest : manifest error : ", manifestErr)
		return
	}

	mjson, _ := json.MarshalIndent(manifest, "", "  ")

	os.WriteFile("goex/exeFiles/manifest.json", mjson, 0644)

}
