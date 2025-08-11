package database

import (
	"wlczak/shokuin/database/schema"
	"wlczak/shokuin/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection struct {
	DB *gorm.DB
}

func GetDB() (Connection, error) {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Europe/Prague"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

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
