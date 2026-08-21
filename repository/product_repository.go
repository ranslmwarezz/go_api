package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductRepository struct {
	conn *sql.DB
}

func NewProductRepository(connection *sql.DB) *ProductRepository {
	return &ProductRepository{conn: connection}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM product"

	rows, err := pr.conn.Query(query)

	if rows.Err() != nil {
		return []model.Product{}, nil
	}

	if err != nil {
		fmt.Println(err)
		return []model.Product{}, nil
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {

		err = rows.Scan(&productObj.ID,
			&productObj.Name,
			&productObj.Price)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, nil
		}

		productList = append(productList, productObj)

	}

	rows.Close()

	return productList, nil

}
