package schema

import "gorm.io/gorm"

type ItemTemplate struct {
	gorm.Model
	Name           string `json:"name"`
	Barcode        string `json:"barcode"`
	Category       int    `json:"category"`
	ExpectedExpiry int    `json:"expected_expiry"`
	Image          string `json:"image"`
}
