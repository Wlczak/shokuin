package api

import (
	"errors"
	"io"
	"net/http"
	"os"
	"slices"

	"github.com/wlczak/shokuin/logger"
	"github.com/wlczak/shokuin/routes/error_handl"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var FilePaths = []string{"item_template"}
var FileTypes = []string{"png"}

func HandleImageUploadApi(c *gin.RouterGroup) {
	c.POST("/upload/:path/:filetype", uploadImage)
	c.GET("/get/:path/:uuid", getImage)
}

func uploadImage(c *gin.Context) {
	zap := logger.GetLogger()
	zap.Info("uploading image")

	path := c.Param("path")
	fileType := c.Param("filetype")
	var imageUuid string

	isUnique := false
	for !isUnique {
		imageUuid = uuid.NewString()

		_, err := os.Stat(path + "/" + imageUuid + "." + fileType)
		if err == nil {
			isUnique = false
		} else if os.IsNotExist(err) {
			isUnique = true
		}
	}

	fullPath := "images" + "/" + path + "/" + imageUuid + "." + fileType

	if !slices.Contains(FilePaths, path) {
		error_handl.HandleErrorJson(c, errors.New("invalid path"))
		return
	}

	if !slices.Contains(FileTypes, fileType) {
		error_handl.HandleErrorJson(c, errors.New("invalid file type"))
		return
	}

	data, err := io.ReadAll(c.Request.Body)

	if err != nil {
		error_handl.HandleErrorJson(c, err)
		zap.Error(err.Error())
		return
	}

	err = os.WriteFile(fullPath, data, 0644)

	if err != nil {
		zap.Error(err.Error())
		error_handl.HandleErrorJson(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"image_uuid": imageUuid,
	})
}
func getImage(c *gin.Context) {
	zap := logger.GetLogger()
	zap.Info("getting image")

	path := c.Param("path")
	uuid := c.Param("uuid")

	fullPath := "images" + "/" + path + "/" + uuid + ".png"

	data, err := os.ReadFile(fullPath)

	if err != nil {
		zap.Error(err.Error())
		error_handl.HandleErrorJson(c, err)
		return
	}

	c.Data(http.StatusOK, "image/png", data)
}
