package core

import (
	"encoding/base64"
	"github.com/rs/xid"
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

type EncodedFilename struct {
	id, filename, encodedName string
}

func (e EncodedFilename) encodeFilename() string {
	uid := xid.New()
	e.id = uid.String()
	rawString := e.filename + "___" + uid.String()
	e.encodedName = base64.StdEncoding.EncodeToString([]byte(rawString))
	return e.encodedName
}

func NewEncodedFilename(filename string) *EncodedFilename {
	e := new(EncodedFilename)
	e.filename = filename
	return e
}
