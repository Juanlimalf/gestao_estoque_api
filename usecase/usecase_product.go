package usecase

import (
	"main/model"
	"main/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func CreateProductUseCase(r repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: r,
	}
}

func (uc *ProductUseCase) GetAllProducts() []model.Product {
	return uc.repository.FetchAllProducts()
}

func (uc *ProductUseCase) CreateProduct(product model.Product) error {
	return uc.repository.InsertProduct(product)
}
