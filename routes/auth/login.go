package auth

import (
	"errors"
	"net/http"
	"wlczak/shokuin/database/model"
	"wlczak/shokuin/database/schema"

	"github.com/gin-gonic/gin"
)

func HandleLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"title": "Login",
	})
}

func HandleLoginPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	var err error = nil
	if username == "" || password == "" {
		err = errors.New("username or password is empty")
	}

	if err != nil {
		c.HTML(http.StatusOK, "login.tmpl", gin.H{
			"title":   "Login",
			"message": err.Error(),
		})
		return
	}

	err = model.CheckUser(&schema.User{Username: username, Password: password})

	if err != nil {
		c.HTML(http.StatusOK, "login.tmpl", gin.H{
			"title":   "Login",
			"message": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "auth_success.tmpl", gin.H{
		"title":   "Login",
		"message": "Successfully logged in",
	})

}
