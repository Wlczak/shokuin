package routes

import (
	"fmt"
	"net/http"
	"strings"
	"wlczak/shokuin/middleware"
	"wlczak/shokuin/routes/api"
	"wlczak/shokuin/routes/auth"
	"wlczak/shokuin/routes/form"
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
		g1.Use(middleware.Auth(utils.AuthLevelNone))
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

	r.Group("/admin", middleware.Auth(utils.AuthLevelAdmin)).GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "admin")
	})

	userGr := r.Group("/user", middleware.Auth(utils.AuthLevelUser))

	userGr.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "user")
	})

	forms := r.Group("/form")
	{
		forms.Use(middleware.Auth(utils.AuthLevelUser))
		forms.GET("/additem", form.HandleAddItem)
	}

	apig := r.Group("/api")
	{
		apig.Match([]string{"GET"}, "/*any", func(c *gin.Context) {
			routes := r.Routes()
			c.Header("Content-Type", "text/html")
			for _, route := range routes {
				if strings.HasPrefix(route.Path, c.Request.RequestURI) && route.Path != "/api/*any" {
					c.String(http.StatusOK, fmt.Sprintf("<a href=\"%s\">%s</a><br>", route.Path, route.Path))
				}
			}
		})
		api.HandleItemApi(apig.Group("/item", middleware.ApiAuth(utils.AuthLevelNone)))

		api.HandleItemTemplateApi(apig.Group("/itemtemplate", middleware.ApiAuth(utils.AuthLevelNone)))

	}

	return r
}
