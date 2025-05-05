package main

import (
	"net/http"
	"wlczak/shokuin/database"
	"wlczak/shokuin/logger"
	"wlczak/shokuin/routing/auth"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	gin.SetMode(gin.DebugMode)

	r := gin.Default()

	r.Use(gin.CustomRecovery(func(c *gin.Context, err any) {
		c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
			"error": err,
		})
	}))

	r.StaticFS("/static", http.Dir("static"))

	r.LoadHTMLGlob("templates/*")

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Home",
		})
	})

	r.GET("/register", auth.HandleRegister)
	r.POST("/register", auth.HandleRegisterPost)

	return r
}

func main() {

	zap := logger.GetLogger()

	zap.Info("Starting server")

	r := setupRouter()

	d, err := database.GetDB()

	if err != nil {
		zap.Error(err.Error())
	}

	d.Setup()

	// Listen and Server in 0.0.0.0:8080
	err = r.Run(":8080")

	if err != nil {
		zap.Error(err.Error())
	}
}
