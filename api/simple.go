package api

import (
	"net/http"
	"os"
)

func SimpleEndpoint(w http.ResponseWriter, r *http.Request) {
	cwd, _ := os.Getwd()
	w.Write([]byte("Simple Endpoint " + cwd))
}
