package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/HasinduLanka/diviyago/pkg/goex"
)

func HelloEndpoint(w http.ResponseWriter, r *http.Request) {

	saveAllErr := goex.DeployEmbedFiles("/tmp/diviyago/exeCache")

	if saveAllErr != nil {
		log.Println("/api/simple : save all files error : ", saveAllErr)
		w.Write([]byte("/api/simple : save all files error : " + saveAllErr.Error()))
		return
	}

	// list all files in current directory

	err := filepath.Walk("/tmp/diviyago/exeCache", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Printf("dir: %v: name: %s\n", info.IsDir(), path)
		fmt.Fprintln(w, info.IsDir(), " -- ", path)

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	w.Write([]byte("HelloEndpoint"))
}
