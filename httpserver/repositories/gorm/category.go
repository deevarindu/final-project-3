package gorm

import (
	"database/sql"

	"github.com/deevarindu/final-project-3/httpserver/repositories"
	"github.com/deevarindu/final-project-3/httpserver/repositories/models"
	"github.com/jinzhu/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) repositories.CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (c *categoryRepository) GetCategories() (*[]models.Category, error) {
	var categories []models.Category
	err := c.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}

	if len(categories) == 0 {
		return nil, sql.ErrNoRows
	}

	return &categories, nil
}
