package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HasinduLanka/diviyago/pkg/goex"
)

func SimpleEndpoint(wr http.ResponseWriter, req *http.Request) {

	saveAllErr := goex.DeployEmbedFiles("/tmp/diviyago/exeCache/")

	if saveAllErr != nil {
		log.Panicln("/api/simple : save all files error : ", saveAllErr)
		wr.Write([]byte("/api/simple : save all files error : " + saveAllErr.Error()))
		return
	}

	AppRun, AppRunErr := goex.ExcecProgramToString("/tmp/diviyago/exeCache/EmbedFiles/ffmpeg-linux-amd64/ffmpeg",
		"-i", "/tmp/diviyago/exeCache/EmbedFiles/ffmpeg-linux-amd64/cloudflare.png", "/tmp/diviyago/exeCache/EmbedFiles/ffmpeg-linux-amd64/cloudflare.webp")

	if AppRunErr != nil {
		log.Panicln("/api/simple : AppRun error : ", AppRunErr)
		wr.Write([]byte("/api/simple : AppRun error : " + AppRunErr.Error()))
		return
	}

	result := AppRun

	wr.Write([]byte("Simple Endpoint " + fmt.Sprint(result)))
}
