package repository

import (
	"fita/model/entity"

	"github.com/stretchr/testify/mock"
)

type CheckoutRepositoryMock struct {
	Mock mock.Mock
}

func(repo *CheckoutRepositoryMock) GetProduct(productId string) (*entity.Products, error) {
	arguments := repo.Mock.Called(productId)
	result := arguments.Get(0).(entity.Products)
	return &result, nil
}

func(repo *CheckoutRepositoryMock) GetPromo(orderId string) (*entity.Promotion, error) {
	arguments := repo.Mock.Called(orderId)
	result := arguments.Get(0).(entity.Promotion)
	return &result,nil
}

func(repo *CheckoutRepositoryMock) Proceed(order *[]entity.Order) error {
	return nil
}

func(repo *CheckoutRepositoryMock) UpdateStock(orderId string, qty uint16) error {
	return nil
}






