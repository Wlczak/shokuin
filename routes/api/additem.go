package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddItem struct {
	Name  string `json:"name"`
	Price int32  `json:"price"`
}

func AddItemApi(c *gin.Context) {

	var additem AddItem
	c.ShouldBindJSON(&additem)

	c.JSON(http.StatusOK, additem)

}
