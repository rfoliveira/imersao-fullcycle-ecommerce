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
	rows, err := pd.db.Query("select id, name, description, price, categoryID, imageURL from products")

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
			&product.ImageURL
		)

		if err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil
}

func (cd *ProductDB) GetProduct(id string) (*entities.Product, error) {
	var product entities.Product
	err := cd.db.QueryRow("select id, name from products where id = ?").Scan(&product.ID, &product.Name)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (cd *ProductDB) CreateProduct(product *entities.Product) (string, error) {
	_, err := cd.db.Exec("insert into products (id, name, description, price, categoryID, imageURL) values (?, ?, ?, ?, ?, ?)", 
		product.ID, 
		product.Name,
		product.Description,
		product.Price,
		product.CategoryID,
		product.ImageURL
	)

	if err != nil {
		return "", err
	}

	return product.ID, nil
}
