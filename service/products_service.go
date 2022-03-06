package service

import (
	"fita/model/entity"
	"fita/repository"
)

type ProductService interface {
	Products() (*[]entity.Products, error)
	Promotion() (*[]entity.Promotion, error)
}

type productServiceImpl struct {
	Repository repository.ProductRepository
}

func NewProductService(
	pr repository.ProductRepository,
) ProductService {
	return &productServiceImpl{pr}
}

func(s *productServiceImpl) Products() (*[]entity.Products, error){
	result, err := s.Repository.FindAllProducts()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func(s *productServiceImpl) Promotion() (*[]entity.Promotion, error){
	result, err := s.Repository.FindAllPromotion()
	if err != nil {
		return nil, err
	}
	return result, nil
}