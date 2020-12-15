package server

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func upload(c *gin.Context) {

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully ", file.Filename))
}

func setupRouter() *gin.Engine {

	// Disable Console Color
	// gin.DisableConsoleColor()
	router := gin.Default()
	router.LoadHTMLGlob("server/template/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload_file.html", gin.H{})
	})
	router.POST("/upload", upload)
	router.StaticFS("/file", http.Dir("server/public"))
	return router
}

func StartServer() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
