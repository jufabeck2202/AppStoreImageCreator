package server

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
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

func CORS() gin.HandlerFunc {
	// TO allow CORS
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
func setupRouter() *gin.Engine {

	// Disable Console Color
	// gin.DisableConsoleColor()
	router := gin.Default()
	store, error := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	if error != nil {
		log.Fatalf("failed to decode: %s", error)
	}
	router.Use(sessions.Sessions("mysession", store))
	router.Use(static.Serve("/", static.LocalFile("../appstoreimagecreatorfrontend/build/", true)))
	router.Use(static.Serve("/create", static.LocalFile("../appstoreimagecreatorfrontend/build/", true)))

	router.POST("/upload", upload)
	router.StaticFS("/file", http.Dir("server/public"))
	router.GET("/test", func(c *gin.Context) {
		session := sessions.Default(c)

		if session.Get("hello") != "world" {
			session.Set("hello", "world")
			session.Save()
			c.JSON(200, gin.H{"hello": session.Get("hello")})
		}else {
			c.JSON(200, gin.H{"ID": session.ID()})
		}

	})
	return router
}

func StartServer() {
	r := setupRouter()

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
