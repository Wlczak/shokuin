package main

import (
	"net/http"
	"wlczak/shokuin/database"

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

	return r
}

func main() {
	r := setupRouter()

	d := database.GetDB()
	d.Setup()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
