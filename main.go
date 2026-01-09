package main

import (
	"fmt"
	"os"
	"time"
	"wlczak/shokuin/database"
	"wlczak/shokuin/logger"
	"wlczak/shokuin/routes"

	_ "wlczak/shokuin/docs"

	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

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

	r := routes.SetupRouter()
	if isProd := os.Getenv("IS_PROD"); isProd == "false" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// Listen and Server in 0.0.0.0:8080
	err = r.Run(":8080")

	if err != nil {
		zap.Error(err.Error())
	}
}
