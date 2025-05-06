package schema

import (
	"wlczak/shokuin/utils"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int             `json:"id"`
	Username  string          `json:"username"`
	Email     string          `json:"email"`
	Password  string          `json:"password"`
	AuthLevel utils.AuthLevel `json:"auth_level"`
}
