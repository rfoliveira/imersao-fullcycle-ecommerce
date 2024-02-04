package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rfoliveira/imersao-fullcycle-ecommerce/internal/database"
	"github.com/rfoliveira/imersao-fullcycle-ecommerce/internal/services"
	"github.com/rfoliveira/imersao-fullcycle-ecommerce/internal/webserver"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/imersao17")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	categoryService := services.NewCategoryService(*categoryDB)

	productDB := database.NewProductDB(db)
	productService := services.NewProductService(*productDB)

	categoryHandler := webserver.NewCategoryHandler(categoryService)
	productHandler := webserver.NewProductHandler(productService)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/category/{id}", categoryHandler.GetCategory)
	router.Get("/category", categoryHandler.GetCategories)
	router.Post("/category", categoryHandler.CreateCategory)

	router.Get("/product/{id}", productHandler.GetProduct)
	router.Get("/product/category/{categoryID}", productHandler.GetProductByCategoryID)
	router.Get("/product", productHandler.GetProducts)
	router.Post("/product", productHandler.CreateProduct)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", router)
}
