package database

import (
	"wlczak/shokuin/database/schema"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
}

func Connect() {

	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Europe/Prague"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&schema.User{})
}
