package entity

import "gorm.io/gorm"

type Promotion struct {
	gorm.Model
	ProductID string `json:"product_id"`
	TypePromo string `json:"type_promo"`
	Amount float64 `json:"amount"`
	Item string `json:"item"`
	Qty uint16 `json:"qty"`
	Percentage uint8 `json:"percentage"`
}