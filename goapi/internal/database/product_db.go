package database

import (
	"database/sql"

	"github.com/rfoliveira/imersao-fullcycle-ecommerce/internal/entities"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{db: db}
}

func (pd *ProductDB) GetProducts() ([]*entities.Product, error) {
	rows, err := pd.db.Query("select id, name, description, price, category_id, image_url from products")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*entities.Product

	for rows.Next() {
		var product entities.Product
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.CategoryID,
			&product.ImageURL,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil
}

func (pd *ProductDB) GetProduct(id string) (*entities.Product, error) {
	var product entities.Product
	err := pd.db.QueryRow("select id, name, description, price, category_id, image_url from products where id = ?", id).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.CategoryID,
		&product.ImageURL,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (pd *ProductDB) GetProductByCategoryID(categoryID string) ([]*entities.Product, error) {
	rows, err := pd.db.Query("select id, name, description, price, category_id, image_url from products where category_id = ?", categoryID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*entities.Product
	for rows.Next() {
		var product entities.Product
		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.CategoryID,
			&product.ImageURL,
		); err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil
}

func (pd *ProductDB) CreateProduct(product *entities.Product) (string, error) {
	_, err := pd.db.Exec("insert into products (id, name, description, price, category_id, image_url) values (?, ?, ?, ?, ?, ?)",
		product.ID,
		product.Name,
		product.Description,
		product.Price,
		product.CategoryID,
		product.ImageURL,
	)

	if err != nil {
		return "", err
	}

	return product.ID, nil
}
