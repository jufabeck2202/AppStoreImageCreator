package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jufabeck2202/AppStoreImageCreator/core"
	"github.com/rs/xid"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

type DataID struct {
	Id     string `json:"id"`
	Device string `json:"device"`
}

type ConvertJob struct {
	Mode           string `json:"mode"`
	BackgroundType string `json:"background"`
	SingleColor    string `json:"color"`
	Gradient1      string `json:"gradient1"`
	Gradient2      string `json:"gradient2"`
}

type JobResult struct {
	URL  string
	name string
}

type JobReply struct {
	results []JobResult
}

func upload(c *gin.Context) {
	id := c.Param("id")
	width, err := strconv.Atoi(c.PostForm("width"))
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	height, err := strconv.Atoi(c.PostForm("height"))
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	filename := filepath.Join("./Storage", "offline", id, filepath.Base(file.Filename))
	if err := c.SaveUploadedFile(file, filename); err != nil {
		print(err.Error())
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}
	frames := core.Frames{}
	frameStruct := frames.GetForSize(width, height)
	c.JSON(http.StatusOK, DataID{Id: id, Device: frameStruct.Name})

}

func firstUpload(c *gin.Context) {
	uid := xid.New()
	core.CreateFolder(uid.String())
	// Source
	width, err := strconv.Atoi(c.PostForm("width"))
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	height, err := strconv.Atoi(c.PostForm("height"))
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	filename := filepath.Join("./Storage", "offline", uid.String(), filepath.Base(file.Filename))
	if err := c.SaveUploadedFile(file, filename); err != nil {
		print(err.Error())
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}
	frames := core.Frames{}
	frameStruct := frames.GetForSize(width, height)
	c.JSON(http.StatusOK, DataID{Id: uid.String(), Device: frameStruct.Name})
}

func process(c *gin.Context) {
	var job ConvertJob
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")

	var frameswg sync.WaitGroup

	files, err := core.FilePathWalkDir(filepath.Join("./Storage", "offline", id))
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	for _, file := range files {
		frameswg.Add(1)
		go core.AddFrame(&frameswg, file, id, job.Gradient1, job.Gradient2)
	}
	//wait until all is finished
	frameswg.Wait()

	results, err := core.FilePathWalkDir(filepath.Join("./Storage", "live", id))
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	response := JobReply{}
	for _, result := range results {
		response.results = append(response.results, JobResult{name: filepath.Base(result), URL: filepath.Join("http://localhost:8080/", strings.Replace(result, "Storage/live/", "converted/", 1))})
	}
	c.JSON(http.StatusOK, response)
}

func file(c *gin.Context) {
	id := c.Param("id")
	filename := c.Param("filename")
	targetPath := filepath.Join("./Storage", "live", id, filename)
	//todo make sure user cant access not allowed files
	if !fileExists(targetPath) {
		c.String(http.StatusBadRequest, fmt.Sprintf("File Not Found"))
		return
	}

	//Seems this headers needed for some browsers (for example without this headers Chrome will download files as txt)
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/octet-stream")
	c.File(targetPath)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
