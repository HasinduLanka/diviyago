package goex

import (
	"embed"
	"io/fs"
	"log"
	"os"
	"strings"
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

	fs.WalkDir(embedcontent, ".", func(path string, dir os.DirEntry, err error) error {
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

			log.Println("saved file : ", saveDir+path)
		}

		return nil
	})

	return nil
}
