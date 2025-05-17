package api

import (
	"net/http"
	"wlczak/shokuin/routes/error_handler"

	"github.com/gin-gonic/gin"
)

type AddItem struct {
	Name  string `json:"name"`
	Price int32  `json:"price"`
}

func AddItemApi(c *gin.Context) {

	var additem AddItem
	err := c.ShouldBindJSON(&additem)
	if err != nil {
		error_handler.WriteErrorJson(c, err)

		return
	}
	c.JSON(http.StatusOK, additem)

}
