package config

import (
	"fmt"

	"github.com/deevarindu/final-project-3/httpserver/repositories/models"
	"github.com/jinzhu/gorm"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "kanbanboard"
)

func CreateConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}

	db.Debug().AutoMigrate(&models.User{}, &models.Task{}, &models.Category{})

	return db, nil
}
