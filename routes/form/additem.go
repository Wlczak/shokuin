package form

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleAddItem(c *gin.Context) {
	c.HTML(http.StatusOK, "additem.tmpl", gin.H{
		"title": "Add Item",
	})
}

func HandleAddItemPost(c *gin.Context) {
	c.HTML(http.StatusOK, "additem.tmpl", gin.H{
		"title": "Add Item",
	})
}
