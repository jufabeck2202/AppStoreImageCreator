package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jufabeck2202/AppStoreImageCreator/core"
	"github.com/rs/xid"
	"net/http"
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
	Mode     string `json:"mode"`
	BackgroundType string `json:"background"`
	SingleColor  string `json:"color"`
	Gradient1 string `json:"gradient1"`
	Gradient2 string `json:"gradient2"`
}

type JobReply struct {
	ResultURLs  []string
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
	CreateFolder(uid.String())
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

	filename := filepath.Join("./Storage","offline", uid.String(), filepath.Base(file.Filename))
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

	files, err := FilePathWalkDir(filepath.Join("./Storage", "offline", id))
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	for _, file := range files {
		frameswg.Add(1)
		go core.AddFrame(&frameswg, file, id, job.Gradient1, job.Gradient2)
	}
	frameswg.Wait()

	results, err := FilePathWalkDir(filepath.Join("./Storage", "live", id))
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	resultPictures := []string{}
	for _, result := range results {
		resultPictures = append(resultPictures, filepath.Join("http://localhost:8080/",strings.Replace(result,"Storage/live/", "converted/",1)))
	}
	println(resultPictures[0])
	c.JSON(http.StatusOK, JobReply{ResultURLs: resultPictures})
}

func file(c *gin.Context) {
	fileName := "result.jpg"
	targetPath := filepath.Join("./", fileName)

	//Seems this headers needed for some browsers (for example without this headers Chrome will download files as txt)
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "application/octet-stream")
	c.File(targetPath)
}
