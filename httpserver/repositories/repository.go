package repositories

import "github.com/deevarindu/final-project-3/httpserver/repositories/models"

type UserRepository interface {
	GetUsers() (*[]models.User, error)
}

type CategoryRepository interface {
	GetCategories() (*[]models.Category, error)
}
