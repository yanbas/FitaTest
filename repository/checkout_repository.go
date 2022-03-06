package repository

import "fita/model/entity"

type CheckoutRepository interface{
	GetProduct(productId string) (*entity.Products, error)
	GetPromo(orderId string) (*entity.Promotion, error)
	Proceed(order *[]entity.Order) error
	UpdateStock(orderId string, qty uint16) error
}