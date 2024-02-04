package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rfoliveira/imersao-fullcycle-ecommerce/internal/entities"
	"github.com/rfoliveira/imersao-fullcycle-ecommerce/internal/services"
)

type CategoryHandler struct {
	CategoryService services.CategoryService
}

func NewCategoryHandler(categoryService *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{CategoryService: *categoryService}
}

func (ch *CategoryHandler) GetCategories(rw http.ResponseWriter, rr *http.Request) {
	categories, err := ch.CategoryService.GetCategories()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(rw).Encode(categories)
}

func (ch *CategoryHandler) GetCategory(rw http.ResponseWriter, rr *http.Request) {
	id := chi.URLParam(rr, "id")
	if id == "" {
		http.Error(rw, "id is required", http.StatusBadRequest)
		return
	}
	category, err := ch.CategoryService.GetCategory(id)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(rw).Encode(category)
}

func (ch *CategoryHandler) CreateCategory(rw http.ResponseWriter, rr *http.Request) {
	var category entities.Category
	err := json.NewDecoder(rr.Body).Decode(&category)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := ch.CategoryService.CreateCategory(category.Name)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(rw).Encode(result)
}
