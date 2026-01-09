package api_schema

import "time"

type AddItem struct {
	Name    string `json:"name"`
	Barcode string `json:"barcode"`
}

type Item struct {
	ID             uint      `json:"id"`
	ItemTemplateId int       `json:"item_template_id"`
	ExpiryDate     time.Time `json:"expiry_date"`
}
