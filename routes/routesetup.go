package routes

import (
	"net/http"
	"wlczak/shokuin/routes/auth"
	"wlczak/shokuin/utils"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

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

	g1 := r.Group("/")
	{
		g1.Use(auth.AuthMiddleware(utils.AuthLevelNone))
		g1.GET("/", func(c *gin.Context) {

			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"title":   "Home",
				"message": "hi",
			})
		})
	}

	r.GET("/register", auth.HandleRegister)
	r.POST("/register", auth.HandleRegisterPost)

	r.GET("/login", auth.HandleLogin)
	r.POST("/login", auth.HandleLoginPost)

	r.Group("/admin", auth.AuthMiddleware(utils.AuthLevelAdmin)).GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "admin")
	})

	r.Group("/user", auth.AuthMiddleware(utils.AuthLevelUser)).GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "user")
	})
	return r
}
