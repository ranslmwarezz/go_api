package repository

import (
	"database/sql"
	"errors"
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

	if err != nil {
		return []model.Product{}, fmt.Errorf("Error retrieving the product: %w", err)
	}

	defer rows.Close()

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {

		err = rows.Scan(&productObj.ID,
			&productObj.Name,
			&productObj.Price)

		if err != nil {
			return []model.Product{}, err
		}

		productList = append(productList, productObj)

	}

	return productList, nil

}

func (pr *ProductRepository) CreateProduct(p model.Product) (int, error) {

	var id int

	query := `INSERT INTO product (product_name, price)
	VALUES ($1, $2)
	RETURNING id`

	err := pr.conn.QueryRow(query, p.Name, p.Price).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("Error to create product: %w", err)
	}

	return id, nil

}

func (pr *ProductRepository) GetProductByID(id_product int) (*model.Product, error) {

	query := `SELECT id, product_name, price
	FROM product
	WHERE id = $1`

	var product model.Product

	err := pr.conn.QueryRow(query, id_product).Scan(&product.ID, &product.Name, &product.Price)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("Error: %w", err)
		}
		return nil, fmt.Errorf("Error trying to find the product: %w", err)
	}

	return &product, nil

}

func (pr *ProductRepository) DeleteById(id_product int) error {

	query := `DELETE FROM product WHERE id = $1`

	row, err := pr.conn.Exec(query, id_product)

	if err != nil {
		return err
	}

	affected, _ := row.RowsAffected()

	if affected == 0 {
		return fmt.Errorf("Product not found")
	}
	return nil

}

func (pr *ProductRepository) UpdateById(id_product int, product model.Product) error {

	query := `UPDATE product 
	SET product_name = $1, price = $2
	WHERE id = $3`

	result, err := pr.conn.Exec(query, product.Name, product.Price, id_product)

	if err != nil {
		return fmt.Errorf("Error updating the product: %w", err)
	}

	affected, err := result.RowsAffected()

	if err != nil {
		return fmt.Errorf("Error checking for updates: %w", err)
	}

	if affected == 0 {
		return fmt.Errorf("Product not found")
	}

	return nil

}
