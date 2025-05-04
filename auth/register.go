package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.tmpl", gin.H{
		"title": "Register",
	})
}

func HandleRegisterPost(c *gin.Context) {

	//username := c.PostForm("username")
	password := c.PostForm("password")
	passwordRepeat := c.PostForm("password-repeat")

	if password != passwordRepeat {
		c.HTML(http.StatusOK, "register.tmpl", gin.H{
			"title":   "Register",
			"message": "Passwords don't match",
		})
		return
	}

	c.HTML(http.StatusOK, "auth_success.tmpl", gin.H{
		"title":   "Register",
		"message": "Registered",
	})
}
