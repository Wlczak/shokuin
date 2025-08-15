package api

import (
	"net/http"
	"wlczak/shokuin/database"
	"wlczak/shokuin/database/model"
	"wlczak/shokuin/database/schema"
	"wlczak/shokuin/logger"
	api_schema "wlczak/shokuin/routes/api/schema"
	"wlczak/shokuin/routes/error_handl"

	"github.com/gin-gonic/gin"
)

func HandleItemTemplateApi(c *gin.RouterGroup) {
	c.POST("/create", AddItemTemplateApi)

	getg := c.Group("/get")
	{
		getg.POST("/by_barcode", GetItemTemplateByBarcodeApi)
	}

}

func AddItemTemplateApi(c *gin.Context) {
	var response api_schema.Response
	var request schema.ItemTemplate

	err := c.ShouldBindJSON(&request)
	if err != nil {
		error_handl.HandleErrorJson(c, err)
		return
	}

	err = model.IsItemTemplateOverlap(&request)

	if err != nil {
		error_handl.HandleErrorJson(c, err)
		return
	}

	db, err := database.GetDB()

	if err != nil {
		zap := logger.GetLogger()
		zap.Error(err.Error())
		panic(err)
	}

	err = db.DB.Create(&request).Error
	if err != nil {
		zap := logger.GetLogger()
		zap.Error(err.Error())
		error_handl.HandleErrorJson(c, err)
		return
	}

	response.Success = true
	response.Message = "Item template added successfully"
	response.Code = http.StatusOK

	c.JSON(response.Code, response)
}

func GetItemTemplateByBarcodeApi(c *gin.Context) {
	var response api_schema.Response
	var request schema.ItemTemplate
	var itemTemplate schema.ItemTemplate

	err := c.ShouldBindJSON(&request)
	if err != nil {
		error_handl.HandleErrorJson(c, err)
		return
	}

	db, err := database.GetDB()

	if err != nil {
		zap := logger.GetLogger()
		zap.Error(err.Error())
		panic(err)
	}

	var count int64
	db.DB.Model(&schema.ItemTemplate{}).Where("barcode = ?", request.Barcode).Count(&count)

	if count == 0 {
		response.Success = false
		response.Message = "Item template not found"
		response.Code = http.StatusNotFound
		c.JSON(response.Code, response)
		return
	} else {
		err = db.DB.Where("barcode = ?", request.Barcode).First(&itemTemplate).Error
		if err != nil {
			zap := logger.GetLogger()
			zap.Error(err.Error())
			error_handl.HandleErrorJson(c, err)
			return
		}
	}

	response.Success = true
	response.Message = "Item template added successfully"
	response.Code = http.StatusOK
	response.Data = itemTemplate

	c.JSON(response.Code, response)
}
