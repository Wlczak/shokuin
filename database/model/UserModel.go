package model

import (
	"wlczak/shokuin/database"
	"wlczak/shokuin/database/schema"
)

func RegisterUser(user *schema.User) {
	db := database.GetDB()
	db.DB.Create(user)
}
