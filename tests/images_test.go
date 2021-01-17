package tests

import (
	"fmt"
	"github.com/jufabeck2202/AppStoreImageCreator/core"
	"github.com/jufabeck2202/AppStoreImageCreator/server"
	"net/http"
	"path/filepath"
	"testing"
)

func TestHelloName(t *testing.T) {
	frames := core.Frames{}

	files, _ := server.FilePathWalkDir(filepath.Join("./inputImages"))

	for _, file := range files {
		frameswg.Add(1)
		go core.AddFrame(&frameswg, file, id, job.Gradient1, job.Gradient2)
	}

}
