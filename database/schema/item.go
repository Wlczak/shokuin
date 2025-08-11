package schema

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ItemTemplateId int       `json:"item_template_id"`
	ExpiryDate     time.Time `json:"expiry_date"`
}
