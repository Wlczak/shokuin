package api

import (
	"net/http"
	"wlczak/shokuin/database"
	"wlczak/shokuin/database/model"
	"wlczak/shokuin/database/schema"
	"wlczak/shokuin/logger"
	api_schema "wlczak/shokuin/routes/api/schema"

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

// GetItemTemplateByBarcodeApi returns the item template with the given barcode
// @Summary Get an item template by barcode
// @Description Returns the item template with the given barcode
// @Tags Item template
// @Accept json
// @Produce json
// @Param barcode path string true "Item template barcode"
// @Success 200 {object} api_schema.ItemTemplate
// @Failure 400 "Invalid request body"
// @Failure 404 "Item template not found"
// @Failure 500 "Internal server error"
// @Router /api/v1/item_template/barcode/{barcode} [get]
func (a *ApiController) GetItemTemplateByBarcodeApi(c *gin.Context) {
	zap := logger.GetLogger()
	barcode := c.Param("barcode")
	if barcode == "" {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	var DBitemTemplate schema.ItemTemplate

	db, err := database.GetDB()

	if err != nil {
		zap.Error(err.Error())
		c.JSON(http.StatusInternalServerError, nil)
	}

	var count int64
	err = db.DB.Model(&schema.ItemTemplate{}).Where("barcode = ?", barcode).Count(&count).Error

	if err != nil {
		zap.Error(err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	if count == 0 {
		c.JSON(http.StatusNotFound, nil)
		return
	} else {
		err = db.DB.Where("barcode = ?", barcode).First(&DBitemTemplate).Error
		if err != nil {
			zap := logger.GetLogger()
			zap.Error(err.Error())
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
	}

	itemTemplate := &api_schema.ItemTemplate{
		Name:           DBitemTemplate.Name,
		Barcode:        DBitemTemplate.Barcode,
		Category:       DBitemTemplate.Category,
		ExpectedExpiry: DBitemTemplate.ExpectedExpiry,
		Image:          DBitemTemplate.Image,
	}

	c.JSON(http.StatusOK, itemTemplate)
}
