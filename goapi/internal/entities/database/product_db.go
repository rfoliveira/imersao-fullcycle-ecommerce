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

func (pd *ProductDB) GetProducts() ([]entities.Product, error) {
	rows, err := pd.db.Query("select id, name, description, price, categoryID, imageURL from products")

}
