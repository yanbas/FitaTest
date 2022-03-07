package repository

import (
	"fita/DB"
	"fita/model/entity"

	"gorm.io/gorm"
)

type productRepo struct {
	
}

func NewProductRepository(conn *gorm.DB) ProductRepository {
	return &productRepo{}
}

func (p *productRepo)FindAllProducts() (*[]entity.Products, error) {
	var products []entity.Products
	err := DB.DB.Find(&products).Error
	if err !=  nil {
		return nil, err
	}
	return &products, nil
}

func (p *productRepo)FindAllPromotion() (*[]entity.Promotion, error) {
	var promotion []entity.Promotion
	err := DB.DB.Find(&promotion).Error
	if err !=  nil {
		return nil, err
	}
	return &promotion, nil
}