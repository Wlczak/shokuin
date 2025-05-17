package api

import (
	"fmt"
	"net/http"
	api_schema "wlczak/shokuin/routes/api/schema"

	"github.com/gin-gonic/gin"
)

func AddItemTemplateApi(c *gin.Context) {

	var itemtemplate api_schema.ItemTemplate
	err := c.ShouldBindJSON(&itemtemplate)
	if err != nil {
		//error_handl.WriteErrorJson(c, err)
		return
	}
	fmt.Println(itemtemplate)
	c.JSON(http.StatusOK, itemtemplate)
}
