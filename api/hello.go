package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func HelloEndpoint(w http.ResponseWriter, r *http.Request) {

	cwd, cwdErr := os.Getwd()

	if cwdErr != nil {
		log.Panicln("/api/hello : cwd error : ", cwdErr)
		w.Write([]byte("/api/hello : cwd error : " + cwdErr.Error()))
		return
	}

	// list all files in current directory

	files, filesErr := os.Open(cwd)

	if filesErr != nil {
		log.Panicln("/api/hello : files error : ", filesErr)
		w.Write([]byte("/api/hello : files error : " + filesErr.Error()))
		return
	}

	defer files.Close()

	filesInfo, filesInfoErr := files.Readdir(-1)

	if filesInfoErr != nil {
		log.Panicln("/api/hello : filesInfo error : ", filesInfoErr)
		w.Write([]byte("/api/hello : filesInfo error : " + filesInfoErr.Error()))
		return
	}

	for _, file := range filesInfo {
		fmt.Fprintln(w, file.Name()+"\n")
	}

	w.Write([]byte("HelloEndpoint"))
}
