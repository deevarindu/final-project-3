package controllers

import (
	"github.com/deevarindu/final-project-3/httpserver/services"
	"github.com/gin-gonic/gin"
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
	ctx.Writer.Header().Set("Content-Type", "application/json")
	response := t.svc.GetTasks()
	WriteJsonResponse(ctx, response)
}
