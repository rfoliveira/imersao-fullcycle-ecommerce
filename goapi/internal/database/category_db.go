package database

import (
	"database/sql"
	"github.com/rfoliveira/imersao-fullcycle-ecommerce/internal/entities"
)

type CategoryDB struct {
	db *sql.DB
}

func NewCategoryDB(db *sql.DB) *CategoryDB {
	return &CategoryDB{db: db}
}

func (cd *CategoryDB) GetCategories() ([]*entities.Category, error) {
	rows, err := cd.db.Query("select id, name from categories")

	if err != nil {
		return nil, err
	}
	// Só pra não precisar colocar no final
	// Isso garante que será executado depois de tudo que estiver
	// depois desse comando
	defer rows.Close()

	var categories []*entities.Category

	for rows.Next() {
		var category entities.Category

		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}

		categories = append(categories, &category)
	}

	return categories, nil
}

func (cd *CategoryDB) GetCategory(id string) (*entities.Category, error) {
	var category entities.Category
	err := cd.db.QueryRow("select id, name from categories where id = ?", id).Scan(&category.ID, &category.Name)
	if err != nil {
		return nil, err
	}

	return &category, nil
}
func (cd *CategoryDB) CreateCategory(category *entities.Category) (string, error) {
	_, err := cd.db.Exec("insert into categories (id, name) values (?, ?)", category.ID, category.Name)

	if err != nil {
		return "", err
	}

	return category.ID, nil
}

func (cd *CategoryDB) UpdateCategory(category *entities.Category) (*entities.Category, error) {
	// todo

	var newCategory entities.Category

	return &newCategory, nil
}

func (cd *CategoryDB) DeleteCategory(id string) (string, error) {
	// todo

	return "", nil
}
