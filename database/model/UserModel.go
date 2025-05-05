package model

import (
	"wlczak/shokuin/database"
	"wlczak/shokuin/database/schema"
	"wlczak/shokuin/logger"
)

func RegisterUser(user *schema.User) {
	db, err := database.GetDB()

	if err != nil {
		zap := logger.GetLogger()
		zap.Error(err.Error())
	}

	db.DB.Create(user)

}
