package error_page

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WriteErrorPage(c *gin.Context, err any) {
	c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
		"error": err,
	})
}
