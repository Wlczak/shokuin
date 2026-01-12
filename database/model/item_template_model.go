package model

import (
	"errors"

	"github.com/wlczak/shokuin/database"
	"github.com/wlczak/shokuin/database/schema"
	"github.com/wlczak/shokuin/logger"
	api_schema "github.com/wlczak/shokuin/routes/api/schema"

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

func IsItemTemplateOverlap(itemTemplate *api_schema.ItemTemplate) error {
	db := getItemTemplateModel()

	var count int64

	db.Where(schema.ItemTemplate{Name: itemTemplate.Name}).Count(&count)

	if count != 0 {
		return errors.New("item template with name " + itemTemplate.Name + " already exists")
	}

	return nil
}
