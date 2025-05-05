package model

import (
	"fmt"
	"wlczak/shokuin/database"
	"wlczak/shokuin/database/schema"
	"wlczak/shokuin/logger"

	"gorm.io/gorm"
)

func getModel() gorm.DB {
	db, err := database.GetDB()

	if err != nil {
		zap := logger.GetLogger()
		zap.Error(err.Error())
		panic(err)
	}

	return *db.DB.Model(&schema.User{})
}

func RegisterUser(user *schema.User) error {
	db := getModel()

	var count int64

	db.Where(schema.User{Username: user.Username}).Count(&count)

	if count != 0 {
		return fmt.Errorf("user with username %s already exists", user.Username)
	}

	db.Create(user)

	return nil
}
