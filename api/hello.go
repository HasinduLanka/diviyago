package api

import (
	"net/http"
)

func HelloEndpoint(wr http.ResponseWriter, req *http.Request) {

	wr.Write([]byte("DiviyaGo - Self contained image processing library for Go"))

}
