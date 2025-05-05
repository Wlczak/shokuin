package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"title": "Login",
	})
}

func HandleLoginPost(c *gin.Context) {
	c.String(http.StatusOK, "WIP")
}
