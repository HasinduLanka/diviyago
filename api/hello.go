package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/HasinduLanka/diviyago/goex"
)

func HelloEndpoint(w http.ResponseWriter, r *http.Request) {

	os.Chdir("/tmp")
	cwd, cwdErr := os.Getwd()

	if cwdErr != nil {
		log.Panicln("/api/hello : cwd error : ", cwdErr)
		w.Write([]byte("/api/hello : cwd error : " + cwdErr.Error()))
		return
	}

	fmt.Fprintln(w, "cwd : ", cwd)

	saveAllErr := goex.SaveAllFiles("exeCache")

	if saveAllErr != nil {
		log.Panicln("/api/simple : save all files error : ", saveAllErr)
		w.Write([]byte("/api/simple : save all files error : " + saveAllErr.Error()))
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
