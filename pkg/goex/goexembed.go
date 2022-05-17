package goex

import (
	"embed"
	"encoding/json"
	"io/fs"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/HasinduLanka/diviyago/pkg/symembed"
)

//go:embed EmbedFiles/*
var embedcontent embed.FS

const fully_written_file = "EmbedFiles/.fully_written_v10001"

func GetEmbededFile(fileName string) ([]byte, error) {
	return embedcontent.ReadFile(fileName)
}

// Save all files in embeded content to the given directory
func DeployEmbedFiles(saveDir string) error {

	if !strings.HasSuffix(saveDir, "/") {
		saveDir = saveDir + "/"
	}

	fullyWritten, fullyWrittenErr := fsexists(saveDir + fully_written_file)

	if fullyWrittenErr != nil {
		return fullyWrittenErr
	}

	if fullyWritten {
		log.Println("DeployEmbedFiles : already fully written")
		return nil
	}

	// Delete the directory to be sure
	os.RemoveAll(saveDir)

	fswg := sync.WaitGroup{}

	fs.WalkDir(embedcontent, ".", func(path string, dir os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if dir.IsDir() {
			// create directory
			os.MkdirAll(saveDir+path, os.ModePerm)
		} else {

			fswg.Add(1)

			go func() {

				defer fswg.Done()
				// create file
				fileBytes, fileErr := embedcontent.ReadFile(path)
				if fileErr != nil {
					return
				}

				fileWrErr := os.WriteFile(saveDir+path, fileBytes, os.ModePerm)
				if fileWrErr != nil {
					return
				}

				// log.Println("saved file : ", saveDir+path)

			}()

		}

		return nil
	})

	manifestBytes, manifestBytesErr := embedcontent.ReadFile("EmbedFiles/manifest.json")

	if manifestBytesErr != nil {
		return manifestBytesErr
	}

	var manifest symembed.SymManifest

	manifestErr := json.Unmarshal(manifestBytes, &manifest)

	if manifestErr != nil {
		return manifestErr
	}

	fswg.Wait()

	applyMnErr := manifest.ApplyManifest(saveDir + "EmbedFiles")

	if applyMnErr != nil {
		return applyMnErr
	}

	// Create the file to be sure
	os.WriteFile(saveDir+fully_written_file, []byte(time.Now().UTC().GoString()), os.ModePerm)

	log.Println("DeployEmbedFiles : wrote all files")

	return nil
}

// exists returns whether the given file or directory exists
func fsexists(path string) (bool, error) {

	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
