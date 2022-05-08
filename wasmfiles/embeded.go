package wasmfiles

import "embed"

//go:embed *
var embedcontent embed.FS

func GetFile(fileName string) ([]byte, error) {
	return embedcontent.ReadFile(fileName)
}
