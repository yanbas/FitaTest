package entity

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	SKU string	`json:"sku"`
	Name string	`json:"name"`
	Price float64	`json:"price"`
	Qty uint16	`json:"qty"`
}