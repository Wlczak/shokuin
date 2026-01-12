package routes

import (
	"context"
	"net/http"
	"time"

	"github.com/wlczak/shokuin/routes/api"
	"golang.org/x/time/rate"

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

	// r.LoadHTMLGlob("templates/*")

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// r.GET("/register", auth.HandleRegister)
	// r.POST("/register", auth.HandleRegisterPost)

	// r.GET("/login", auth.HandleLogin)
	// r.POST("/login", auth.HandleLoginPost)

	// r.Group("/admin", middleware.Auth(utils.AuthLevelAdmin)).GET("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "admin")
	// })

	// userGr := r.Group("/user", middleware.Auth(utils.AuthLevelUser))

	// userGr.GET("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "user")
	// })

	// forms := r.Group("/form")
	// {
	// 	forms.Use(middleware.Auth(utils.AuthLevelUser))
	// 	forms.GET("/additem", form.HandleAddItem)
	// }

	//	@title			Shokuin API
	//	@version		1.0

	//	@host		localhost:8080
	//	@BasePath	/api/v1

	const apiRequestLimit = 5
	const apiRequestInterval = time.Minute / 10 // 10 requests per minute
	const apiRequestTimeout = time.Second * 30

	apiv1 := r.Group("/api/v1")
	{
		openFoodFactsApiLimiter := rate.NewLimiter(rate.Every(apiRequestInterval), apiRequestLimit)

		c := api.ApiController{}
		item := apiv1.Group("/item")
		{
			item.GET(":id", c.GetItemApi)
			item.POST("", c.AddItemApi)
			item.DELETE(":id", c.DeleteItemApi)
			item.PATCH(":id", c.PatchItemApi)
		}

		item_template := apiv1.Group("/item_template")
		{
			item_template.GET(":id", c.GetItemTemplateApi)
			item_template.GET("/barcode/:barcode", c.GetItemTemplateByBarcodeApi)
			item_template.POST("", c.AddItemTemplateApi)
			item_template.DELETE(":id", c.DeleteItemTemplateApi)
			item_template.PATCH(":id", c.PatchItemTemplateApi)

			item_template.GET("/open_food_facts/:barcode", func(gin *gin.Context) {
				ctx, cancel := context.WithTimeout(context.Background(), apiRequestTimeout)
				defer cancel()

				err := openFoodFactsApiLimiter.Wait(ctx)
				if err != nil {
					gin.JSON(http.StatusTooManyRequests, nil)
					return
				}
				c.GetItemTemplateFromOpenFoodFacts(gin)
			})
		}
	}

	return r
}
