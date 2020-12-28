package server

import (
	"os"
	"path/filepath"
)

func CreateFolder(id string) {
	newpath := filepath.Join(".", "Storage", id)
	os.MkdirAll(newpath, os.ModePerm)
}

func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}