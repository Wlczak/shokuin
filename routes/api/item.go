package api

import (
	"net/http"
	"time"

	"wlczak/shokuin/database"
	"wlczak/shokuin/database/schema"
	"wlczak/shokuin/logger"
	api_schema "wlczak/shokuin/routes/api/schema"
	"wlczak/shokuin/routes/error_handl"

	"github.com/gin-gonic/gin"
)

// GetItemApi returns the item with the given id
// @Summary Get an item
// @Description Returns the item with the given id
// @Tags item
// @Accept json
// @Produce json
// @Param id path string true "Item ID"
// @Success 200 {object} api_schema.Item
// @Failure 404 "Item not found"
// @Router /item/{id} [get]
func (a *ApiController) GetItemApi(c *gin.Context) {
	id := c.Param("id")

	db, err := database.GetDB()

	if err != nil {
		zap := logger.GetLogger()
		zap.Error(err.Error())
		panic(err)
	}

	var dbitem schema.Item
	db.DB.Where("id = ?", id).First(&dbitem)

	if dbitem.ID == 0 {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	item := api_schema.Item{
		ID:             dbitem.ID,
		ItemTemplateId: dbitem.ItemTemplateId,
		ExpiryDate:     dbitem.ExpiryDate,
	}
	c.JSON(http.StatusOK, item)
}

// AddItemApi adds a new item to the database
// @Summary Add a new item
// @Description Adds a new item. The expiry date must not be older than 30 days.
// @Tags item
// @Accept json
// @Produce json
// @Param item body api_schema.AddItem true "Item to add"
// @Success 204 {string} string "No Content"
// @Failure 400 "Invalid request body"
// @Failure 500 "Internal server error"
// @Router /item [post]
func (a *ApiController) AddItemApi(c *gin.Context) {
	var request schema.Item

	err := c.ShouldBindJSON(&request)
	if err != nil {
		error_handl.HandleErrorJson(c, err)
		return
	}

	if request.ExpiryDate.Before(time.Now().Add(-time.Hour * 24 * 30)) {
		return
	}

	// @todo add template id check

	db, err := database.GetDB()

	if err != nil {
		zap := logger.GetLogger()
		zap.Error(err.Error())
		panic(err)
	}

	db.DB.Create(&request)

	c.JSON(http.StatusNoContent, nil)
}
