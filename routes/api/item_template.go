package api

import (
	"net/http"
	"wlczak/shokuin/database"
	"wlczak/shokuin/database/model"
	"wlczak/shokuin/database/schema"
	"wlczak/shokuin/logger"
	api_schema "wlczak/shokuin/routes/api/schema"
	"wlczak/shokuin/routes/error_handl"

	"github.com/gin-gonic/gin"
)

// GetItemTemplateApi returns the item template with the given id
// @Summary Get an item template
// @Description Returns the item template with the given id
// @Tags Item template
// @Accept json
// @Produce json
// @Param id path string true "Item template ID"
// @Success 200 {object} api_schema.ItemTemplate
// @Failure 400 "Invalid request body"
// @Failure 404 "Item template not found"
// @Failure 500 "Internal server error"
// @Router /api/v1/item_template/{id} [get]
func (a *ApiController) GetItemTemplateApi(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	db, err := database.GetDB()

	if err != nil {
		zap := logger.GetLogger()
		zap.Error(err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	var dbitem schema.ItemTemplate
	db.DB.Where("id = ?", id).First(&dbitem)

	if dbitem.ID == 0 {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, dbitem)

}

// AddItemTemplateApi adds a new item template to the database
// @Summary Add a new item template
// @Description Adds a new item template
// @Tags Item template
// @Accept json
// @Produce json
// @Param item_template body api_schema.ItemTemplate true "Item template to add"
// @Success 204 "Succesfully added (no content)"
// @Failure 304 "Item template with name already exists"
// @Failure 400 "Invalid request body"
// @Failure 500 "Internal server error"
// @Router /api/v1/item_template [post]
func (a *ApiController) AddItemTemplateApi(c *gin.Context) {
	var request api_schema.ItemTemplate
	zap := logger.GetLogger()
	err := c.ShouldBindJSON(&request)
	if err != nil {
		zap.Error(err.Error())
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	err = model.IsItemTemplateOverlap(&request)

	if err != nil {
		zap.Error(err.Error())
		c.JSON(http.StatusNotModified, &api_schema.ErrorMessage{
			Error: err.Error(),
		})
		return
	}

	db, err := database.GetDB()

	if err != nil {
		zap.Error(err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	err = db.DB.Create(&request).Error
	if err != nil {
		zap.Error(err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// DeleteItemTemplateApi deletes the item template with the given id
// @Summary Delete an item template
// @Description Deletes the item template with the given id
// @Tags Item template
// @Accept json
// @Produce json
// @Param id path string true "Item template ID"
// @Success 204 "Succesfully deleted (no content)"
// @Failure 400 "Invalid request body"
// @Failure 404 "Item template not found"
// @Router /api/v1/item_template/{id} [delete]
func (a *ApiController) DeleteItemTemplateApi(c *gin.Context) {
	id := c.Param("id")
	zap := logger.GetLogger()

	if id == "" {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	db, err := database.GetDB()

	if err != nil {
		zap.Error(err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	err = db.DB.Where("id = ?", id).Delete(&schema.ItemTemplate{}).Error
	if err != nil {
		zap.Error(err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// PatchItemTemplateApi updates the item template with the given id
// @Summary Update an item template
// @Description Updates the item template with the given id
// @Tags Item template
// @Accept json
// @Produce json
// @Param id path string true "Item template ID"
// @Param item_template body api_schema.ItemTemplate true "Item template to update"
// @Success 204 "Succesfully updated (no content)"
// @Failure 400 "Invalid request body"
// @Failure 404 "Item template not found"
// @Router /api/v1/item_template/{id} [patch]
func (a *ApiController) PatchItemTemplateApi(c *gin.Context) {
	id := c.Param("id")
	zap := logger.GetLogger()

	if id == "" {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	db, err := database.GetDB()

	if err != nil {
		zap.Error(err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	var request api_schema.ItemTemplate
	err = c.ShouldBindJSON(&request)
	if err != nil {
		zap.Error(err.Error())
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	err = db.DB.Model(&schema.ItemTemplate{}).Where("id = ?", id).Updates(request).Error
	if err != nil {
		zap.Error(err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func GetItemTemplateByBarcodeApi(c *gin.Context) {
	var response api_schema.Response
	var request schema.ItemTemplate
	var itemTemplate schema.ItemTemplate

	err := c.ShouldBindJSON(&request)
	if err != nil {
		error_handl.HandleErrorJson(c, err)
		return
	}

	db, err := database.GetDB()

	if err != nil {
		zap := logger.GetLogger()
		zap.Error(err.Error())
		panic(err)
	}

	var count int64
	db.DB.Model(&schema.ItemTemplate{}).Where("barcode = ?", request.Barcode).Count(&count)

	if count == 0 {
		response.Success = false
		response.Message = "Item template not found"
		response.Code = http.StatusNotFound
		c.JSON(response.Code, response)
		return
	} else {
		err = db.DB.Where("barcode = ?", request.Barcode).First(&itemTemplate).Error
		if err != nil {
			zap := logger.GetLogger()
			zap.Error(err.Error())
			error_handl.HandleErrorJson(c, err)
			return
		}
	}

	response.Success = true
	response.Message = "Item template added successfully"
	response.Code = http.StatusOK
	response.Data = itemTemplate

	c.JSON(response.Code, response)
}
