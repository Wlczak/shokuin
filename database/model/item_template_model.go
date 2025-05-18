package model

import (
	"errors"
	"wlczak/shokuin/database"
	"wlczak/shokuin/database/schema"
	"wlczak/shokuin/logger"

	"gorm.io/gorm"
)

func getItemTemplateModel() gorm.DB {

	db, err := database.GetDB()

	if err != nil {
		zap := logger.GetLogger()
		zap.Error(err.Error())
		panic(err)
	}

	return *db.DB.Model(&schema.ItemTemplate{})
}

func IsItemTemplateOverlap(itemTemplate *schema.ItemTemplate) error {
	db := getItemTemplateModel()

	var count int64

	db.Where(schema.ItemTemplate{Name: itemTemplate.Name}).Count(&count)

	if count != 0 {
		return errors.New("item template with name " + itemTemplate.Name + " already exists")
	}

	return nil
}
