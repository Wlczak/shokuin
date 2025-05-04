package database

import (
	"wlczak/shokuin/database/schema"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection struct {
	DB *gorm.DB
}

func GetDB() Connection {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Europe/Prague"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return Connection{DB: db}
}

func (d Connection) Setup() {
	if d.DB == nil {
		panic("db isn't initialized")
	}

	d.DB.AutoMigrate(&schema.User{})
}
