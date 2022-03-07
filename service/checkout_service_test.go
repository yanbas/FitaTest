package service

import (
	"fita/model/entity"
	"fita/repository"
	"fmt"
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Register Repository
var checkoutRepository = &repository.CheckoutRepositoryMock{Mock: mock.Mock{}}
// Register Services
var checkoutSvc = NewCheckoutService(checkoutRepository)

func TestCheckoutMacbookPro(t *testing.T) {
	products := []entity.Products{
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

	product := entity.Products{
			SKU: "43N23P",
			Name: "Macbook Pro",
			Price: 5399.99,
			Qty: 5,
	}

	promo := entity.Promotion{
			ProductID: "43N23P",
			TypePromo: "item",
			Item: "120P90",
			Amount: 0,
			Qty: 1,
	}

	productId := "43N23P"

	param := graphql.ResolveParams{
		Args: map[string]interface{}{
			"ID_PRO":"43N23P",
			"QTY":"1",
		},
	}

	checkoutRepository.Mock.On("GetProduct", productId).Return(product).Once()
	checkoutRepository.Mock.On("GetPromo", productId).Return(promo).Once()
	checkoutRepository.Mock.On("GetProduct", promo.Item).Return(products[0]).Once()

	result, _ := checkoutSvc.Checkout(param)

	resultStruct := result.([]entity.Order)

	assert.NotNil(t, result, "Cannot nil")
	assert.Equal(t, float64(0), resultStruct[1].GrandTotal, "Grand Total Promo must be 0")
	assert.Equal(t, 2, len(resultStruct), "Total record must be 2")
	assert.Equal(t, "43N23P", resultStruct[0].ProductID, "Product ID is 43N23P")
	assert.Equal(t, "120P90", resultStruct[1].ProductID, "Product Promo ID is 43N23P")	
}


func TestCheckoutThreeGoogleHome(t *testing.T) {
	product := entity.Products{
			SKU: "120P90",
			Name: "Google Home",
			Price: 49.99,
			Qty: 10,
	}

	promo := entity.Promotion{
			ProductID: "120P90",
			TypePromo: "amount",
			Item: "",
			Amount: 49.99,
			Qty: 3,
	}

	productId := "120P90"

	param := graphql.ResolveParams{
		Args: map[string]interface{}{
			"ID_PRO":"120P90",
			"QTY":"3",
		},
	}

	checkoutRepository.Mock.On("GetProduct", productId).Return(product).Once()
	checkoutRepository.Mock.On("GetPromo", productId).Return(promo).Once()

	result, _ := checkoutSvc.Checkout(param)

	resultStruct := result.([]entity.Order)

	assert.NotNil(t, result, "Cannot nil")
	assert.Equal(t, float64(99.97999999999999), resultStruct[0].GrandTotal, fmt.Sprintf("Grand Total must be %.2f",float64(99.97999999999999)))
	assert.Equal(t, "120P90", resultStruct[0].ProductID, "Product ID is 120P90")
}


func TestCheckoutDiscountPercentage(t *testing.T) {
	product := entity.Products{
			SKU: "A304SD",
			Name: "Alexa Speaker",
			Price: 109.50,
			Qty: 10,
	}

	promo := entity.Promotion{
			ProductID: "A304SD",
			TypePromo: "percentage",
			Item: "",
			Amount: 0,
			Qty: 3,
			Percentage: 10,
	}

	productId := "A304SD"

	param := graphql.ResolveParams{
		Args: map[string]interface{}{
			"ID_PRO":"A304SD",
			"QTY":"3",
		},
	}

	checkoutRepository.Mock.On("GetProduct", productId).Return(product).Once()
	checkoutRepository.Mock.On("GetPromo", productId).Return(promo).Once()

	result, _ := checkoutSvc.Checkout(param)

	resultStruct := result.([]entity.Order)

	assert.NotNil(t, result, "Cannot nil")
	assert.Equal(t, float64((product.Price * 3) - ((product.Price * 3) * float64(promo.Percentage) /100)), resultStruct[0].GrandTotal, fmt.Sprintf("Grand Total must be %.2f",float64((product.Price * 3) - ((product.Price * 3) * float64(promo.Percentage) /100))))
	assert.Equal(t, "A304SD", resultStruct[0].ProductID, "Product ID is A304SD")
}