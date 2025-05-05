package auth

import (
	"errors"
	"net/http"
	"os"
	"time"
	"wlczak/shokuin/database/model"
	"wlczak/shokuin/database/schema"
	"wlczak/shokuin/logger"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"password": password,
		"time":     time.Now().Unix(),
	}).SignedString([]byte(os.Getenv("APP_KEY")))

	if err != nil {
		zap := logger.GetLogger()
		zap.Error(err.Error())
		return
	}

	c.SetCookie("SHOKUIN_JWT", token, 3600, "", "", false, true)
	c.HTML(http.StatusOK, "auth_success.tmpl", gin.H{
		"title":   "Login",
		"message": "Successfully logged in",
	})

}
