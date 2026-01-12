package model

import (
	"fmt"

	"github.com/wlczak/shokuin/database"
	"github.com/wlczak/shokuin/database/schema"
	"github.com/wlczak/shokuin/logger"
	"github.com/wlczak/shokuin/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func getUserModel() gorm.DB {
	db, err := database.GetDB()

	if err != nil {
		zap := logger.GetLogger()
		zap.Error(err.Error())
		panic(err)
	}

	return *db.DB.Model(&schema.User{})
}

func RegisterUser(user *schema.User) error {
	db := getUserModel()

	var count int64

	db.Where(schema.User{Username: user.Username}).Count(&count)

	if count != 0 {
		return fmt.Errorf("user with username %s already exists", user.Username)
	}

	db.Create(user)

	return nil
}

// TODO: Separete logic and database calls - logic should be in the routes handler
func CheckUser(user *schema.User) error {
	db := getUserModel()

	var count int64
	var DBUser schema.User

	db.Where(schema.User{Username: user.Username}).Count(&count)

	if count == 0 {
		return fmt.Errorf("user with username %s doesn't exist %d ", user.Username, count)
	}

	db.Where(schema.User{Username: user.Username}).First(&DBUser)

	err := bcrypt.CompareHashAndPassword([]byte(DBUser.Password), []byte(user.Password))

	if err != nil {
		return fmt.Errorf("wrong password")
	}

	return nil
}

func GetUserLevelByUsername(user *schema.User) utils.AuthLevel {
	db := getUserModel()

	var DBUser schema.User

	db.Where(schema.User{Username: user.Username}).First(&DBUser)

	return DBUser.AuthLevel
}
