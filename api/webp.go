package api

import (
	"log"
	"net/http"
	"os"

	"github.com/HasinduLanka/diviyago/goex"
)

func WebpEndpoint(wr http.ResponseWriter, req *http.Request) {

	os.Chdir("/tmp")

	saveAllErr := goex.SaveAllFiles("exeCache")

	if saveAllErr != nil {
		log.Panicln("/api/webp : save all files error : ", saveAllErr)
		wr.Write([]byte("/api/webp : save all files error : " + saveAllErr.Error()))
		return
	}

	_, AppRunErr := goex.ExcecProgramToString("/tmp/exeCache/exeFiles/ffmpeg-linux-amd64/ffmpeg",
		"-i", "/tmp/exeCache/exeFiles/ffmpeg-linux-amd64/cloudflare.png", "/tmp/exeCache/exeFiles/ffmpeg-linux-amd64/cloudflare.webp")

	if AppRunErr != nil {
		log.Panicln("/api/simple : AppRun error : ", AppRunErr)
		wr.Write([]byte("/api/simple : AppRun error : " + AppRunErr.Error()))
		return
	}

	outfile := "/tmp/exeCache/exeFiles/ffmpeg-linux-amd64/cloudflare.webp"

	// write file to response

	file, err := os.ReadFile(outfile)
	if err != nil {
		log.Panicln("/api/webp : open file error : ", err)
		wr.Write([]byte("/api/webp : open file error : " + err.Error()))
		return
	}

	wr.Header().Set("Content-Type", "image/webp")
	wr.Write(file)

}
