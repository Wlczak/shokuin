package database

import (
	"wlczak/shokuin/database/schema"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type connection struct {
	db *gorm.DB
}

func GetDB() connection {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Europe/Prague"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return connection{db: db}
}

func (d connection) Setup() {
	if d.db == nil {
		panic("db isn't initialized")
	}

	d.db.AutoMigrate(&schema.User{})
}
