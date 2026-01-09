package api

import (
	"net/http"
	"strconv"
	"time"

	"wlczak/shokuin/database"
	"wlczak/shokuin/database/schema"
	"wlczak/shokuin/logger"
	api_schema "wlczak/shokuin/routes/api/schema"

	"github.com/gin-gonic/gin"
)

// GetItemApi returns the item with the given id
// @Summary Get an item
// @Description Returns the item with the given id
// @Tags Item
// @Accept json
// @Produce json
// @Param id path string true "Item ID"
// @Success 200 {object} api_schema.Item
// @Failure 400 "Invalid request body"
// @Failure 404 "Item not found"
// @Failure 500 "Internal server error"
// @Router /item/{id} [get]
func (a *ApiController) GetItemApi(c *gin.Context) {
	zap := logger.GetLogger()
	id := c.Param("id")
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

	var dbitem schema.Item
	err = db.DB.Where("id = ?", id).First(&dbitem).Error

	if err != nil {
		zap.Error(err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	if dbitem.ID == 0 {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	item := api_schema.Item{
		ItemTemplateId: dbitem.ItemTemplateId,
		ExpiryDate:     dbitem.ExpiryDate,
	}
	c.JSON(http.StatusOK, item)
}

// AddItemApi adds a new item to the database
// @Summary Add a new item
// @Description Adds a new item. The expiry date must not be older than 30 days.
// @Tags Item
// @Accept json
// @Produce json
// @Param item body api_schema.Item true "Item to add"
// @Success 204 {string} string "No Content"
// @Failure 400 "Invalid request body"
// @Failure 500 "Internal server error"
// @Router /item [post]
func (a *ApiController) AddItemApi(c *gin.Context) {
	var request schema.Item
	zap := logger.GetLogger()
	err := c.ShouldBindJSON(&request)
	if err != nil {
		zap.Error(err.Error())
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	if request.ExpiryDate.Before(time.Now().Add(-time.Hour * 24 * 30)) {
		c.JSON(http.StatusBadRequest, &api_schema.ErrorMessage{Error: "bad expiry date"})
		return
	}

	// @todo add template id check

	db, err := database.GetDB()

	if err != nil {
		zap := logger.GetLogger()
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

// DeleteItemApi deletes the item with the given id
// @Summary Delete an item
// @Description Deletes the item with the given id
// @Tags Item
// @Accept json
// @Produce json
// @Param id path string true "Item ID"
// @Success 204 "Succesfully deleted (no content)"
// @Failure 400 "Invalid request body"
// @Failure 404 "Item not found"
// @Router /item/{id} [delete]
func (a *ApiController) DeleteItemApi(c *gin.Context) {
	stringId := c.Param("id")
	zap := logger.GetLogger()
	if stringId == "" {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	id, err := strconv.Atoi(stringId)

	if err != nil {
		c.JSON(http.StatusNotModified, nil)
		zap.Error(err.Error())
		return
	}

	db, err := database.GetDB()

	if err != nil {
		zap.Error(err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	var dbitem = &schema.Item{}

	err = db.DB.Where("id = ?", id).First(dbitem).Error

	if err != nil {
		zap.Error(err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	if dbitem.ID == 0 {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	err = db.DB.Delete(dbitem).Error

	if err != nil {
		zap.Error(err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// PatchItemApi updates the item with the given id
// @Summary Update an item
// @Description Updates the item with the given id
// @Tags Item
// @Accept json
// @Produce json
// @Param id path string true "Item ID"
// @Param item body api_schema.Item true "Item to update"
// @Success 204 "Succesfully updated (no content)"
// @Failure 400 "Invalid request body"
// @Failure 404 "Item not found"
// @Router /item/{id} [patch]
func (a *ApiController) PatchItemApi(c *gin.Context) {
	id := c.Param("id")
	zap := logger.GetLogger()

	if id == "" {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	db, err := database.GetDB()

	if err != nil {
		zap := logger.GetLogger()
		zap.Error(err.Error())
		c.JSON(http.StatusBadRequest, nil)
	}

	var dbitem schema.Item
	err = db.DB.Where("id = ?", id).First(&dbitem).Error

	if err != nil {
		zap.Error(err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	if dbitem.ID == 0 {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	var item api_schema.Item
	err = c.ShouldBindJSON(&item)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	dbitem.ItemTemplateId = item.ItemTemplateId
	dbitem.ExpiryDate = item.ExpiryDate

	err = db.DB.Save(&dbitem).Error

	if err != nil {
		zap.Error(err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
