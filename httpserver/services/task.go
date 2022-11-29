package services

import (
	"database/sql"

	"github.com/deevarindu/final-project-3/httpserver/repositories"
	"github.com/deevarindu/final-project-3/httpserver/repositories/models"
	"github.com/deevarindu/final-project-3/httpserver/views"
)

type TaskSvc struct {
	repo repositories.TaskRepository
}

func NewTaskSvc(repo repositories.TaskRepository) *TaskSvc {
	return &TaskSvc{
		repo: repo,
	}
}

func (t *TaskSvc) GetTasks() *views.Response {
	tasks, err := t.repo.GetTasks()
	if err != nil {
		if err == sql.ErrNoRows {
			return views.DataNotFoundResponse(err)
		}
		return views.InternalServerErrorResponse(err)
	}

	return views.SuccessGetResponse(parseModelToGetTasks(tasks), "Success get all tasks")
}

func parseModelToGetTasks(mod *[]models.Task) *[]views.GetTasks {
	var t []views.GetTasks
	for _, v := range *mod {
		t = append(t, views.GetTasks{
			ID:          *v.ID,
			Title:       v.Title,
			Status:      v.Status,
			Description: v.Description,
			UserID:      v.UserID,
			CategoryID:  v.CategoryID,
		})
	}
	return &t
}
