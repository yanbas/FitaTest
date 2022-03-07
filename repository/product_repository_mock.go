package repository

import (
	"fita/model/entity"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func(repo *ProductRepositoryMock) FindAllProducts() (*[]entity.Products, error) {
	arguments := repo.Mock.Called()
	result := arguments.Get(0).([]entity.Products)
	return &result,nil
}

func(repo *ProductRepositoryMock) FindAllPromotion() (*[]entity.Promotion, error) {
	arguments := repo.Mock.Called()
	result := arguments.Get(0).([]entity.Promotion)
	return &result,nil
}
