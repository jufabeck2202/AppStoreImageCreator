package server

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func upload(c *gin.Context) {
	session := sessions.Default(c)
	id :=session.Get("id")

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	if idStr, ok := id.(string); ok {
		print(idStr)
		filename := filepath.Join("./Storage","test",filepath.Base(file.Filename))
		if err := c.SaveUploadedFile(file, filename); err != nil {
			print(err.Error())
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
	} else {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully ", file.Filename))
}