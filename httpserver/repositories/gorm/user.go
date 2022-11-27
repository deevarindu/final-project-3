package gorm

import (
	"database/sql"

	"github.com/deevarindu/final-project-3/httpserver/repositories"
	"github.com/deevarindu/final-project-3/httpserver/repositories/models"
	"github.com/jinzhu/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) GetUsers() (*[]models.User, error) {
	var users []models.User
	err := u.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, sql.ErrNoRows
	}

	return &users, nil
}
