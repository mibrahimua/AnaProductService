package service

import (
	"AnaProductService/model"
	"AnaProductService/repository"
)

type ProductService struct {
	productRepository *repository.ProductRepository
}

func NewUserService(userRepository *repository.ProductRepository) *ProductService {
	return &ProductService{userRepository}
}

func (us *ProductService) GetListProductByName(productName string) ([]model.Product, error) {
	return us.productRepository.GetListProductByName(productName)
}
