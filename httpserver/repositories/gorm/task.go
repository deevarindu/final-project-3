package gorm

import (
	"database/sql"

	"github.com/deevarindu/final-project-3/httpserver/repositories"
	"github.com/deevarindu/final-project-3/httpserver/repositories/models"
	"github.com/jinzhu/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) repositories.TaskRepository {
	return &taskRepository{}
}

func (t *taskRepository) GetTasks() (*[]models.Task, error) {
	var tasks []models.Task
	err := t.db.Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	if len(tasks) == 0 {
		return nil, sql.ErrNoRows
	}

	return &tasks, nil
}
