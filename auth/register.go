package auth

import (
	"net/http"
	"wlczak/shokuin/database/model"
	"wlczak/shokuin/database/schema"

	"github.com/gin-gonic/gin"
)

func HandleRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.tmpl", gin.H{
		"title": "Register",
	})
}

func HandleRegisterPost(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")
	passwordRepeat := c.PostForm("password-repeat")

	if password != passwordRepeat {
		c.HTML(http.StatusOK, "register.tmpl", gin.H{
			"title":   "Register",
			"message": "Passwords don't match",
		})
		return
	}

	model.RegisterUser(&schema.User{Username: username, Email: "test@test.test", Password: password})

	c.HTML(http.StatusOK, "auth_success.tmpl", gin.H{
		"title":   "Register",
		"message": "Registered",
	})
}
