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
	return &taskRepository{
		db: db,
	}
}

func (t *taskRepository) GetTasks() (*[]models.Task, error) {
	var tasks []models.Task
	err := t.db.Preload("User").Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	if len(tasks) == 0 {
		return nil, sql.ErrNoRows
	}

	return &tasks, nil
}

func (t *taskRepository) CreateTask(task *models.Task) error {
	err := t.db.Create(task).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *taskRepository) UpdateTask(task *models.Task) error {
	err := t.db.Save(task).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *taskRepository) DeleteTask(id int) error {
	err := t.db.Delete(&models.Task{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
