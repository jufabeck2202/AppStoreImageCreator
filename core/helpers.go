package core

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/rs/xid"
	"os"
	"path/filepath"
	"strings"
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
		if !info.IsDir() && (filepath.Ext(path) == ".png" || filepath.Ext(path) == ".jpeg" || filepath.Ext(path) == ".jpg") {
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

func (e EncodedFilename) decodeFilename() (string, string, error) {

	data, err := base64.StdEncoding.DecodeString(e.encodedName)

	if err != nil {
		return "", "", errors.New("couldn't parse string")
	}
	idAndFilename := strings.SplitN(string(data), "___", 2)
	filename := idAndFilename[0]
	id := idAndFilename[1]
	fileExtension := filepath.Ext(idAndFilename[0])
	if fileExtension != ".png" && fileExtension != ".jpeg" && fileExtension != ".jpg" {
		return "", "", errors.New("wrong extension")
	}

	fmt.Printf("%q\n", data)
	return filename, id, nil

}

func NewEncodedFilename(filename string) *EncodedFilename {
	e := new(EncodedFilename)
	e.filename = filename
	return e
}
