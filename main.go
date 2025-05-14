package main

import (
	"os"
	"wlczak/shokuin/database"
	"wlczak/shokuin/logger"
	"wlczak/shokuin/routes"

	"github.com/joho/godotenv"
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

	d, err := database.GetDB()

	if err != nil {
		zap.Error(err.Error())
	}

	d.Setup()
	r := routes.SetupRouter()

	// Listen and Server in 0.0.0.0:8080
	err = r.Run(":8080")

	if err != nil {
		zap.Error(err.Error())
	}
}
