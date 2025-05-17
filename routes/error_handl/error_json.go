package error_handl

import (
	"net/http"
	"os"
	"wlczak/shokuin/logger"

	"github.com/gin-gonic/gin"
)

func WriteErrorJson(c *gin.Context, err error) {
	zap := logger.GetLogger()
	zap.Error(err.Error())

	if os.Getenv("IS_PROD") == "false" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "oops, something went wrong",
		})
	}
}
