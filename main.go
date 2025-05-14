package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"wlczak/shokuin/database"
	"wlczak/shokuin/logger"
	"wlczak/shokuin/routes/auth"
	"wlczak/shokuin/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

// var EnvVars = []string{
// 	"DB_HOST",
// 	"DB_PORT",
// 	"DB_USER",
// 	"DB_PASSWORD",
// 	"DB_NAME",
// }

func setupEnv() error {
	_, err := os.Stat(".env")
	if os.IsNotExist(err) {
		f, err := os.Create(".env")
		if err != nil {
			return err
		}
		defer func() {
			if err := f.Close(); err != nil {
				zap := logger.GetLogger()
				zap.Error(err.Error())
			}
		}()
	}

	// it would be nice to finish this, but it desn't have priority rn

	// envMap, err := godotenv.Read(".env")

	// godotenv.Read(".env")
	// for _, value := range EnvVars {
	// 	if envMap[value] == "" {
	// 		godotenv.Write(map[string]string{
	// 			value: "",
	// 		}, ".env")
	// 	}
	// 	fmt.Println(value)
	// }

	err = godotenv.Load()

	return err
}

func main() {
	zap := logger.GetLogger()
	zap.Info("Starting server")

	err := setupEnv()

	if err != nil {
		zap.Fatal(err.Error())
	}

	for {
		d, err := database.GetDB()
		if err != nil {
			zap.Error(err.Error())
			fmt.Println("DB didn't connect properly - retrying in 5s")
			time.Sleep(time.Second * 5)
		} else {
			d.Setup()
			break
		}
	}

	r := setupRouter()

	// Listen and Server in 0.0.0.0:8080
	err = r.Run(":8080")

	if err != nil {
		zap.Error(err.Error())
	}
}
