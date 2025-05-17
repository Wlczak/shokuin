package api_schema

type ItemTemplate struct {
	Name           string `json:"name"`
	Barcode        string `json:"barcode"`
	Category       int    `json:"category"`
	ExpectedExpiry int    `json:"expected_expiry"`
	Image          string `json:"image"`
}
