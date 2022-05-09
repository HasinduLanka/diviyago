package api

import (
	"log"
	"net/http"
	"os"

	"github.com/HasinduLanka/diviyago/goex"
)

func WebpEndpoint(wr http.ResponseWriter, req *http.Request) {

	saveAllErr := goex.SaveAllFiles(`/tmp/diviyago/exeCache/`)

	if saveAllErr != nil {
		log.Panicln(`/api/webp : save all files error : `, saveAllErr)
		wr.Write([]byte(`/api/webp : save all files error : ` + saveAllErr.Error()))
		return
	}

	fileBytes, fileBytesErr := os.ReadFile(`/tmp/diviyago/exeCache/exeFiles/ffmpeg-linux-amd64/cloudflare.png`)

	if fileBytesErr != nil {
		log.Panicln(`/api/webp : file read error : `, fileBytesErr)
		wr.Write([]byte(`/api/webp : file read error : ` + fileBytesErr.Error()))
		return
	}

	file, AppRunErr := goex.ExcecTask(nil, fileBytes, `/tmp/diviyago/exeCache/exeFiles/ffmpeg-linux-amd64/ffmpeg`, `-y`, `-f`, `image2pipe`,
		`-i`, `pipe:`, `-vf`, `scale=360:-1`, `-f`, `webp`, `pipe:`)

	if AppRunErr != nil {
		log.Panicln(`/api/simple : AppRun error : `, AppRunErr)
		wr.Write([]byte(`/api/simple : AppRun error : ` + AppRunErr.Error()))
		return
	}

	wr.Header().Set(`Content-Type`, `image/webp`)
	wr.Write(file)

}
