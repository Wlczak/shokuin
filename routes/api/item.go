package api

import (
	"net/http"

	api_schema "wlczak/shokuin/routes/api/schema"
	"wlczak/shokuin/routes/error_handl"

	"github.com/gin-gonic/gin"
)

func AddItemApi(c *gin.Context) {

	var additem api_schema.AddItem
	err := c.ShouldBindJSON(&additem)
	if err != nil {
		error_handl.HandleErrorJson(c, err)

		return
	}
	c.JSON(http.StatusOK, additem)

}
