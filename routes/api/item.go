package api

import (
	"net/http"
	"time"

	"wlczak/shokuin/database"
	"wlczak/shokuin/database/schema"
	"wlczak/shokuin/logger"
	api_schema "wlczak/shokuin/routes/api/schema"
	"wlczak/shokuin/routes/error_handl"

	"github.com/gin-gonic/gin"
)

func HandleItemApi(c *gin.RouterGroup) {
	c.POST("/create", AddItemApi)
}

func AddItemApi(c *gin.Context) {
	var response api_schema.Response
	var request schema.Item

	err := c.ShouldBindJSON(&request)
	if err != nil {
		error_handl.HandleErrorJson(c, err)
		return
	}

	if request.ExpiryDate.Before(time.Now().Add(-time.Hour * 24 * 30)) {
		response.Success = false
		response.Message = "expiry date is broken"
		c.JSON(response.Code, response)
		return
	}

	// @todo add template id check

	db, err := database.GetDB()

	db.DB.Create(&request)

	if err != nil {
		zap := logger.GetLogger()
		zap.Error(err.Error())
		panic(err)
	}
	response.Success = true
	response.Message = "success"
	response.Code = http.StatusOK
	c.JSON(response.Code, response)
}
