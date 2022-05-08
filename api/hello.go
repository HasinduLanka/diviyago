package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func HelloEndpoint(w http.ResponseWriter, r *http.Request) {

	cwd, cwdErr := os.Getwd()

	if cwdErr != nil {
		log.Panicln("/api/hello : cwd error : ", cwdErr)
		w.Write([]byte("/api/hello : cwd error : " + cwdErr.Error()))
		return
	}

	// list all files in current directory

	err := filepath.Walk(cwd, func(path string, info os.FileInfo, err error) error {
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
