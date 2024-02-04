package services

import (
	"github.com/rfoliveira/imersao-fullcycle-ecommerce/internal/database"
	"github.com/rfoliveira/imersao-fullcycle-ecommerce/internal/entities"
)

type ProductService struct {
	ProductDB database.ProductDB
}

func NewProductService(productDB database.ProductDB) *ProductService {
	return &ProductService{ProductDB: productDB}
}

func (ps *ProductService) GetProducts() ([]*entities.Product, error) {
	products, err := ps.ProductDB.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) GetProduct(id string) (*entities.Product, error) {
	product, err := ps.ProductDB.GetProduct(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *ProductService) GetProductByCategoryID(categoryID string) ([]*entities.Product, error) {
	products, err := ps.ProductDB.GetProductByCategoryID(categoryID)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) CreateProduct(name string, description string, price float64, categoryID string, imageURL string) (*entities.Product, error) {
	product := entities.NewProduct(name, description, price, categoryID, imageURL)
	_, err := ps.ProductDB.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}
