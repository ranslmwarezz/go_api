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

func (pu *ProductUseCase) CreateProduct(product model.Product) (model.Product, error) {

	productID, err := pu.repo.CreateProduct(product)

	if err != nil {
		return model.Product{}, err
	}

	product.ID = productID

	return product, nil

}

func (pr *ProductUseCase) GetProductByID(product_id int) (*model.Product, error) {

	product, err := pr.repo.GetProductByID(product_id)

	if err != nil {
		return nil, err
	}

	return product, nil

}

func (pr *ProductUseCase) DeleteById(product_id int) error {

	err := pr.repo.DeleteById(product_id)

	if err != nil {
		return err
	}

	return nil
}

func (pr *ProductUseCase) UpdateById(product_id int, product model.Product) error {

	err := pr.repo.UpdateById(product_id, product)

	if err != nil {
		return err
	}

	return nil

}
