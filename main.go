package main

import (
	"net/http"
	"time"
	"wlczak/shokuin/auth"
	"wlczak/shokuin/database"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	gin.SetMode(gin.DebugMode)

	r := gin.Default()

	r.Use(gin.CustomRecovery(func(c *gin.Context, err any) {
		sugar := zap.NewExample().Sugar()
		defer sugar.Sync()
		sugar.Infow("failed to fetch URL",
			"url", "http://example.com",
			"attempt", 3,
			"backoff", time.Second,
		)
		sugar.Infof("failed to fetch URL: %s", "http://example.com")

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
	r := setupRouter()

	d := database.GetDB()

	d.Setup()

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
