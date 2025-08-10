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
}

func AddItemTemplateApi(c *gin.Context) {
	var response api_schema.Response
	var request api_schema.ItemTemplate

	err := c.ShouldBindJSON(&request)
	if err != nil {
		error_handl.HandleErrorJson(c, err)
		return
	}

	var itemTemplateSchema schema.ItemTemplate
	err = model.IsItemTemplateOverlap(&itemTemplateSchema)

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

	templateModel := &schema.ItemTemplate{
		Name:     request.Name,
		Barcode:  request.Barcode,
		Category: request.Category,
		Image:    request.Image,
	}

	err = db.DB.Create(&templateModel).Error
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
