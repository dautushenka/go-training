package repository

import (
	"database/sql"
	"go-training/proj/core/model"
)

type CategoryRepository interface {
	GetCategoryList() (*[]model.Category, error)
}

type categoryRepository struct {
	db *sql.DB
}

func CreateCategoryRepository(db *sql.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) GetCategoryList() (*[]model.Category, error) {
	var categories []model.Category
	rows, err := r.db.Query("SELECT Id, Name from categories")
	if err != nil {
		return &categories, err
	}
	defer rows.Close()

	for rows.Next() {
		var cat model.Category
		err = rows.Scan(&cat.Id, &cat.Name)
		if err != nil {
			return &categories, err
		}
		categories = append(categories, cat)
	}

	return &categories, nil
}
