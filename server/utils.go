package server

import (
	"os"
	"path/filepath"
)

func CreateFolder(id string) {
	newpath := filepath.Join(".", "Storage", id)
	os.MkdirAll(newpath, os.ModePerm)
}
