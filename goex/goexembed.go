package goex

import (
	"embed"
	"encoding/json"
	"os"
	"strings"

	"github.com/HasinduLanka/diviyago/pkg/symembed"
)

//go:embed exeFiles/*
var embedcontent embed.FS

func GetFile(fileName string) ([]byte, error) {
	return embedcontent.ReadFile(fileName)
}

// Save all files in embeded content to the given directory
func SaveAllFiles(saveDir string) error {

	if !strings.HasSuffix(saveDir, "/") {
		saveDir = saveDir + "/"
	}

	// Delete the directory if it exists
	os.RemoveAll(saveDir)

	Walk(embedcontent, ".", func(path string, dir os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if dir.IsDir() {
			// create directory
			os.MkdirAll(saveDir+path, os.ModePerm)
		} else {
			// create file
			fileBytes, fileErr := embedcontent.ReadFile(path)
			if fileErr != nil {
				return fileErr
			}

			fileWrErr := os.WriteFile(saveDir+path, fileBytes, os.ModePerm)
			if fileWrErr != nil {
				return fileWrErr
			}

			// log.Println("saved file : ", saveDir+path)
		}

		return nil
	})

	manifestBytes, manifestBytesErr := embedcontent.ReadFile("exeFiles/manifest.json")

	if manifestBytesErr != nil {
		return manifestBytesErr
	}

	var manifest symembed.SymManifest

	manifestErr := json.Unmarshal(manifestBytes, &manifest)

	if manifestErr != nil {
		return manifestErr
	}

	applyMnErr := manifest.ApplyManifest(saveDir + "exeFiles")

	if applyMnErr != nil {
		return applyMnErr
	}

	return nil
}
