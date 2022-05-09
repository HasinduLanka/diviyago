package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/HasinduLanka/diviyago/goex"
)

func SimpleEndpoint(wr http.ResponseWriter, req *http.Request) {

	os.Chdir("/tmp")

	saveAllErr := goex.SaveAllFiles("exeCache")

	if saveAllErr != nil {
		log.Panicln("/api/simple : save all files error : ", saveAllErr)
		wr.Write([]byte("/api/simple : save all files error : " + saveAllErr.Error()))
		return
	}

	AppRun, AppRunErr := goex.ExcecProgramToString("/tmp/exeCache/exeFiles/ffmpeg-linux-amd64/ffmpeg",
		"-i", "/tmp/exeCache/exeFiles/ffmpeg-linux-amd64/cloudflare.png", "/tmp/exeCache/exeFiles/ffmpeg-linux-amd64/cloudflare.webp")

	if AppRunErr != nil {
		log.Panicln("/api/simple : AppRun error : ", AppRunErr)
		wr.Write([]byte("/api/simple : AppRun error : " + AppRunErr.Error()))
		return
	}

	result := AppRun

	wr.Write([]byte("Simple Endpoint " + fmt.Sprint(result)))
}
