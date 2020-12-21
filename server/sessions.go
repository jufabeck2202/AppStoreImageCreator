package server

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"io/ioutil"
	"os"
	"time"
)

func Session() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		id :=session.Get("id")
		if id == nil{
			uid := xid.New()
			session.Set("id", uid.String())
			session.Save()
		} else {
			if str, ok := id.(string); ok {
				if session.Get("createdFolder") != "True" {
					print("not true")
					session.Set("createdFolder", "True")
					session.Save()
					go CreateFolder(str)
				}
			}

		}

	}
}


func isOlderThanOneDay(t time.Time) bool {
	return time.Now().Sub(t) > 24*time.Hour
}

func findFilesOlderThanOneDay(dir string) (files []os.FileInfo, err error) {
	tmpfiles, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	for _, file := range tmpfiles {
		if file.Mode().IsRegular() {
			if isOlderThanOneDay(file.ModTime()) {
				files = append(files, file)
			}
		}
	}
	return
}