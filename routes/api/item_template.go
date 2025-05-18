package api

import (
	"net/http"
	"wlczak/shokuin/database/model"
	"wlczak/shokuin/database/schema"
	api_schema "wlczak/shokuin/routes/api/schema"
	"wlczak/shokuin/routes/error_handl"

	"github.com/gin-gonic/gin"
)

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

	response.Success = true
	response.Message = "Item template added successfully"
	c.JSON(http.StatusOK, response)
}
