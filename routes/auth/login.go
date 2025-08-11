package auth

import (
	"errors"
	"net/http"
	"os"
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

	jwt_exp := os.Getenv("JWT_EXPIRATION")
	jwt_dur, err := time.ParseDuration(jwt_exp + "h")
	if err != nil {
		zap := logger.GetLogger()
		zap.Error(err.Error())
		jwt_dur = time.Hour
	}
	exp := time.Now().Add(jwt_dur)

	token, err := utils.GenToken(jwt.MapClaims{
		"username":   username,
		"time":       time.Now().Unix(),
		"exp":        exp.Unix(),
		"auth_level": authLevel,
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
