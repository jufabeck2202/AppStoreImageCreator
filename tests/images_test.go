package tests


import (
	"fmt"
	"net/http"
	"path/filepath"
	"testing"
	"github.com/jufabeck2202/AppStoreImageCreator/core"
	"github.com/jufabeck2202/AppStoreImageCreator/server"

)

func TestHelloName(t *testing.T) {
	frames := core.Frames{}

	files, _ := server.FilePathWalkDir(filepath.Join("./inputImages"))

	for _, file := range files {
		frameswg.Add(1)
		go core.AddFrame(&frameswg, file, id, job.Gradient1, job.Gradient2)
	}

}