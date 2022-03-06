package repository

import (
	"fita/DB"
	"fita/model/entity"

	"gorm.io/gorm"
)


type checkoutRepo struct {
	connection *gorm.DB
}

func NewCheckoutRepository(conn *gorm.DB) CheckoutRepository {
	return &checkoutRepo{conn}
}


func (c *checkoutRepo)GetProduct(productId string) (*entity.Products, error) {
	var product entity.Products
	err := DB.DB.First(&product).Where("sku",productId).Error
	if err !=  nil {
		return nil, err
	}

	return &product, nil
}

func (p *checkoutRepo)GetPromo(orderId string) (*entity.Promotion, error) {
	var promotion entity.Promotion
	err := DB.DB.First(&promotion).Where("product_id",orderId).Error
	if err !=  nil {
		return nil, err
	}
	return &promotion, nil
}

func (p *checkoutRepo)Proceed(order *[]entity.Order) error {
	err := DB.DB.Create(&order).Error
	if err !=  nil {
		return err
	}
	return nil
}

func (p *checkoutRepo)UpdateStock(orderId string, qty uint16) error {
	DB.DB.UpdateColumn("qty", gorm.Expr("qty  - ?", qty))
	return nil
}


