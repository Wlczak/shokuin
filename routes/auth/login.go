package auth

import (
	"errors"
	"net/http"
	"time"
	"wlczak/shokuin/database/model"
	"wlczak/shokuin/database/schema"
	"wlczak/shokuin/logger"
	"wlczak/shokuin/utils"

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

	authLevel := model.GetUserLevelByUsername(&schema.User{Username: username})

	token, err := utils.GenToken(jwt.MapClaims{
		"username":  username,
		"time":      time.Now().Unix(),
		"exp":       time.Now().Add(time.Hour).Unix(),
		"authLevel": authLevel,
	})

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
