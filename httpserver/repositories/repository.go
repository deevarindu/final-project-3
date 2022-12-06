package repositories

import "github.com/deevarindu/final-project-3/httpserver/repositories/models"

type UserRepository interface {
	GetUsers() (*[]models.User, error)
	Register(user *models.User) error
	UpdateUser(user *models.User) error
	DeleteUser(user *models.User) error
}

type CategoryRepository interface {
	GetCategories() (*[]models.Category, error)
	CreateCategory(category *models.Category) error
	UpdateCategory(category *models.Category) error
	DeleteCategory(id int) error
}

type TaskRepository interface {
	GetTasks() (*[]models.Task, error)
	CreateTask(task *models.Task) error
	UpdateTask(task *models.Task) error
	DeleteTask(id int) error
}
