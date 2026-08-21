package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUseCase struct {
	// Repository
	repo repository.ProductRepository
}

func NewProductUseCase(r repository.ProductRepository) *ProductUseCase {
	return &ProductUseCase{repo: r}
}

func (p *ProductUseCase) GetProducts() ([]model.Product, error) {

	return p.repo.GetProducts()

}
