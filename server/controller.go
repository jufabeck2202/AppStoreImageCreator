package server

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("id", 12090292)
	session.Set("email", "test@gmail.com")
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"message": "User Sign In successfully",
	})
}
