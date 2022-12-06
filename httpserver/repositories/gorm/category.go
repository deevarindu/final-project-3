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
	err := c.db.Preload("Tasks").Find(&categories).Error
	if err != nil {
		return nil, err
	}

	if len(categories) == 0 {
		return nil, sql.ErrNoRows
	}

	return &categories, nil
}

func (c *categoryRepository) CreateCategory(category *models.Category) error {
	err := c.db.Create(category).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *categoryRepository) UpdateCategory(category *models.Category) error {
	err := c.db.Save(category).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *categoryRepository) DeleteCategory(id int) error {
	err := c.db.Delete(&models.Category{ID: &id}).Error
	if err != nil {
		return err
	}
	return nil
}
