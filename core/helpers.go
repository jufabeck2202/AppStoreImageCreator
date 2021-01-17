package core

import (
	"os"
	"path/filepath"
)

func CreateFolder(id string) {
	newpath := filepath.Join(".", "Storage", "offline", id)
	newpath2 := filepath.Join(".", "Storage", "live", id)
	os.MkdirAll(newpath, os.ModePerm)
	os.MkdirAll(newpath2, os.ModePerm)

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