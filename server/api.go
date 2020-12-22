package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"net/http"
	"path/filepath"
)

type DataID struct {
	Id string `json:"id"`
}

func upload(c *gin.Context) {
	id := c.Param("id")

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	filename := filepath.Join("./Storage", id, filepath.Base(file.Filename))
	if err := c.SaveUploadedFile(file, filename); err != nil {
		print(err.Error())
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}
	c.JSON(http.StatusOK, DataID{Id: id})

}

func firstUpload(c *gin.Context) {
	uid := xid.New()
	CreateFolder(uid.String())
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	filename := filepath.Join("./Storage", uid.String(), filepath.Base(file.Filename))
	if err := c.SaveUploadedFile(file, filename); err != nil {
		print(err.Error())
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}
	c.JSON(http.StatusOK, DataID{Id: uid.String()})
}
