package controllers

import (
	"net/http"
	"strconv"

	"github.com/deevarindu/final-project-3/httpserver/controllers/params"
	"github.com/deevarindu/final-project-3/httpserver/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TaskController struct {
	svc *services.TaskSvc
}

func NewTaskController(svc *services.TaskSvc) *TaskController {
	return &TaskController{
		svc: svc,
	}
}

func (t *TaskController) GetTasks(ctx *gin.Context) {
	response := t.svc.GetTasks()
	WriteJsonResponse(ctx, response)
}

func (t *TaskController) CreateTask(ctx *gin.Context) {
	var req params.TaskCreateRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = validator.New().Struct(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := t.svc.CreateTask(&req)
	WriteJsonResponse(ctx, response)
}

func (t *TaskController) UpdateTitleDesc(ctx *gin.Context) {
	var req params.TaskUpdateTitleDescRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = validator.New().Struct(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskID, _ := strconv.Atoi(ctx.Param("taskId"))
	response := t.svc.UpdateTitleDesc(&req, taskID)
	WriteJsonResponse(ctx, response)
}

func (t *TaskController) UpdateStatus(ctx *gin.Context) {
	var req params.TaskUpdateStatusRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = validator.New().Struct(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskID, _ := strconv.Atoi(ctx.Param("taskId"))
	response := t.svc.UpdateStatus(&req, taskID)
	WriteJsonResponse(ctx, response)
}

func (t *TaskController) UpdateCategory(ctx *gin.Context) {
	var req params.TaskUpdateCategoryRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = validator.New().Struct(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskID, _ := strconv.Atoi(ctx.Param("taskId"))
	response := t.svc.UpdateCategory(&req, taskID)
	WriteJsonResponse(ctx, response)
}

func (t *TaskController) DeleteTask(ctx *gin.Context) {
	taskID, _ := strconv.Atoi(ctx.Param("taskId"))
	response := t.svc.DeleteTask(taskID)
	WriteJsonResponse(ctx, response)
}
