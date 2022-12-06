package services

import (
	"database/sql"

	"github.com/deevarindu/final-project-3/helper/jwt"
	"github.com/deevarindu/final-project-3/httpserver/controllers/params"
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
			User: views.GetUsers{
				ID:       *v.User.ID,
				FullName: v.User.FullName,
				Email:    v.User.Email,
			},
		})
	}
	return &t
}

func (t *TaskSvc) CreateTask(req *params.TaskCreateRequest) *views.Response {
	task := parseRequestToModelTask(req)
	task.UserID = jwt.UserData.ID

	id := 1
	if tasks, err := t.repo.GetTasks(); err == nil {
		id = len(*tasks) + 1
	}
	task.ID = &id
	task.Status = false

	err := t.repo.CreateTask(task)

	if err != nil {
		return views.InternalServerErrorResponse(err)
	}

	return views.SuccessCreateResponse(task, "Success create task")
}

func parseRequestToModelTask(req *params.TaskCreateRequest) *models.Task {
	return &models.Task{
		Title:       req.Title,
		Description: req.Description,
		CategoryID:  req.CategoryID,
	}
}

func (t *TaskSvc) FindTaskById(id int) *models.Task {
	tasks, _ := t.repo.GetTasks()
	for _, task := range *tasks {
		if *task.ID == id {
			return &task
		}
	}
	return nil
}

func (t *TaskSvc) UpdateTitleDesc(req *params.TaskUpdateTitleDescRequest, id int) *views.Response {
	task := t.FindTaskById(id)
	if task == nil {
		return views.DataNotFoundResponse(nil)
	}

	task.Title = req.Title
	task.Description = req.Description

	err := t.repo.UpdateTask(task)
	if err != nil {
		return views.InternalServerErrorResponse(err)
	}

	return views.SuccessUpdateResponse(task, "Success update task")
}

func (t *TaskSvc) UpdateStatus(req *params.TaskUpdateStatusRequest, id int) *views.Response {
	task := t.FindTaskById(id)
	if task == nil {
		return views.DataNotFoundResponse(nil)
	}

	task.Status = req.Status

	err := t.repo.UpdateTask(task)
	if err != nil {
		return views.InternalServerErrorResponse(err)
	}

	return views.SuccessUpdateResponse(task, "Success update task")
}

func (t *TaskSvc) UpdateCategory(req *params.TaskUpdateCategoryRequest, id int) *views.Response {
	task := t.FindTaskById(id)
	if task == nil {
		return views.DataNotFoundResponse(nil)
	}

	task.CategoryID = req.CategoryID

	err := t.repo.UpdateTask(task)
	if err != nil {
		return views.InternalServerErrorResponse(err)
	}

	return views.SuccessUpdateResponse(task, "Success update task")
}

func (t *TaskSvc) DeleteTask(id int) *views.Response {
	err := t.repo.DeleteTask(id)
	if err != nil {
		return views.InternalServerErrorResponse(err)
	}

	return views.SuccessDeleteResponse("Success delete task")
}
