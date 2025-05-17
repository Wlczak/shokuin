package auth

import (
	"errors"
	"net/http"
	"wlczak/shokuin/database/model"
	"wlczak/shokuin/database/schema"
	error_page "wlczak/shokuin/routes/error_handler"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

	var err error = nil
	if password != passwordRepeat {
		err = errors.New("passwords don't match")
	}

	if username == "" || password == "" || passwordRepeat == "" {
		err = errors.New("username or password is empty")
	}

	if err != nil {
		c.HTML(http.StatusOK, "register.tmpl", gin.H{
			"title":   "Register",
			"message": err.Error(),
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		error_page.WriteErrorPage(c, err)
		return
	}

	err = model.RegisterUser(&schema.User{Username: username, Email: "", Password: string(hash)})

	if err != nil {
		c.HTML(http.StatusOK, "register.tmpl", gin.H{
			"title":   "Register",
			"message": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "auth_success.tmpl", gin.H{
		"title":   "Register",
		"message": "Registered successfully",
	})
}
