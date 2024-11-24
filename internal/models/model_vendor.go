package models

type Vendor struct {
	VendorID string `gorm: "primaryKey" json: "vendorId"`
	Name     string `json: "name"`
	Contact  string `json: "contact"`
	Email    string `json: "email"`
	Phone    string `json: "phone"`
	Address  string `json: "address"`
}
