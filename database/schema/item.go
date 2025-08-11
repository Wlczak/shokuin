package schema

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ItemTemplateId int    `json:"item_template_id"`
	ExpiryDate     string `json:"expiry_date"`
}
