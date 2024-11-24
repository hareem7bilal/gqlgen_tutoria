package models

type Product struct {
	ProductID string  `gorm: "primaryKey" json: "productId"`
	VendorID  string  `json: "vendorId"`
	Name      string  `json: "name"`
	Price     float32 `gorm: "type:numeric" json: "price"`
}
