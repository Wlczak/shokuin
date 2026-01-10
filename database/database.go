package database

import (
	"github.com/wlczak/shokuin/database/schema"
	"github.com/wlczak/shokuin/logger"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Connection struct {
	DB *gorm.DB
}

func GetDB() (Connection, error) {

	db, err := gorm.Open(sqlite.Open("shokuin.db"), &gorm.Config{})

	if err != nil {
		return Connection{DB: nil}, err
	}

	return Connection{DB: db}, nil
}

func (d Connection) Setup() {
	if d.DB == nil {
		zap := logger.GetLogger()
		zap.Panic("db isn't initialized")
		panic("db isn't initialized")
	}

	err := d.DB.AutoMigrate(&schema.User{})

	if err != nil {
		zap := logger.GetLogger()
		zap.Error(err.Error())
		panic(err)
	}

	err = d.DB.AutoMigrate(&schema.ItemTemplate{})

	if err != nil {
		zap := logger.GetLogger()
		zap.Error(err.Error())
		panic(err)
	}

	err = d.DB.AutoMigrate(&schema.Item{})

	if err != nil {
		zap := logger.GetLogger()
		zap.Error(err.Error())
		panic(err)
	}
}
