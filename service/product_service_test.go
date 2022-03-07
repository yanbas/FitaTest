package service

import (
	"fita/model/entity"
	"fita/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Register Repository
var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
// Register Services
var productSvc = NewProductService(productRepository)


func TestGetProduct(t *testing.T) {
	product := []entity.Products{
		{
			SKU: "120P90",
			Name: "Google Home",
			Price: 49.99,
			Qty: 10,
		},
		{
			SKU: "43N23P",
			Name: "Macbook Pro",
			Price: 5399.99,
			Qty: 5,
		},
		{
			SKU: "A304SD",
			Name: "Alexa Speaker",
			Price: 109.50,
			Qty: 10,
		},
		{
			SKU: "120P90",
			Name: "Raspberry Pi B",
			Price: 30.00,
			Qty: 2,
		},	
	}

	productRepository.Mock.On("FindAllProducts").Return(product)
	result, _ := productSvc.Products()

	assert.NotNil(t, result, "Cannot nil")
	assert.Equal(t, 4, len(*result), "Len Data must be 4")
}


func TestGetPromotion(t *testing.T) {
	promo := []entity.Promotion{
		{
			ProductID: "43N23P",
			TypePromo: "item",
			Item: "120P90",
			Amount: 0,
			Qty: 1,
		},
		{
			ProductID: "120P90",
			TypePromo: "amount",
			Item: "",
			Amount: 49.99,
			Qty: 3,
		},
		{
			ProductID: "A304SD",
			TypePromo: "percentage",
			Item: "",
			Amount: 0,
			Qty: 3,
			Percentage: 10,
		},
	}

	productRepository.Mock.On("FindAllPromotion").Return(promo)
	result, _ := productSvc.Promotion()

	assert.NotNil(t, result, "Cannot nil")
	assert.Equal(t, 3, len(*result), "Len Data must be 3")
}