package entity

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderID string `json:"order_id"`
	ProductID string	`json:"product_id"`
	Price float64	`json:"name"`
	PromoAmount float64	`json:"promo"`
	GrandTotal float64	`json:"grand_total"`
	PromoID uint	`json:"promo_id"`
	Qty uint16	`json:"qty"`
}