package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rfoliveira/imersao-fullcycle-ecommerce/internal/entities"
	"github.com/rfoliveira/imersao-fullcycle-ecommerce/internal/services"
)

type ProductHandler struct {
	ProductService services.ProductService
}

func NewProductHandler(productService *services.ProductService) *ProductHandler {
	return &ProductHandler{ProductService: *productService}
}

func (ph *ProductHandler) GetProducts(rw http.ResponseWriter, rr *http.Request) {
	products, err := ph.ProductService.GetProducts()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(rw).Encode(products)
}

func (ph *ProductHandler) GetProduct(rw http.ResponseWriter, rr *http.Request) {
	id := chi.URLParam(rr, "id")
	if id == "" {
		http.Error(rw, "id is required", http.StatusBadRequest)
		return
	}
	product, err := ph.ProductService.GetProduct(id)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(rw).Encode(product)
}

func (ph *ProductHandler) GetProductByCategoryID(rw http.ResponseWriter, rr *http.Request) {
	categoryID := chi.URLParam(rr, "categoryID")
	if categoryID == "" {
		http.Error(rw, "categoryID is required", http.StatusBadRequest)
		return
	}
	products, err := ph.ProductService.GetProductByCategoryID(categoryID)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(rw).Encode(products)
}

func (ph *ProductHandler) CreateProduct(rw http.ResponseWriter, rr *http.Request) {
	var product entities.Product
	err := json.NewDecoder(rr.Body).Decode(&product)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := ph.ProductService.CreateProduct(
		product.Name,
		product.Description,
		product.Price,
		product.CategoryID,
		product.ImageURL,
	)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(rw).Encode(result)
}
