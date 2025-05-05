package main

import (
	"fmt"
	"net/http"
	"os"
	"wlczak/shokuin/database"
	"wlczak/shokuin/logger"
	"wlczak/shokuin/routes/auth"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

	r.GET("/", func(c *gin.Context) {

		tokenString, err := c.Cookie("SHOKUIN_JWT")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("APP_KEY")), nil
		})

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		fmt.Print(token.Claims)

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":   "Home",
			"message": "hi",
		})
	})

	r.GET("/register", auth.HandleRegister)
	r.POST("/register", auth.HandleRegisterPost)

	r.GET("/login", auth.HandleLogin)
	r.POST("/login", auth.HandleLoginPost)

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
		defer f.Close()
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
