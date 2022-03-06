package repository

import "fita/model/entity"

type ProductRepository interface{
	FindAllProducts() (*[]entity.Products, error)
	FindAllPromotion() (*[]entity.Promotion, error)
}