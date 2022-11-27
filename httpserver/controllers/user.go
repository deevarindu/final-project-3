package controllers

import (
	"github.com/deevarindu/final-project-3/httpserver/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	svc *services.UserSvc
}

func NewUserController(svc *services.UserSvc) *UserController {
	return &UserController{
		svc: svc,
	}
}

func (u *UserController) GetUsers(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	response := u.svc.GetUsers()
	WriteJsonResponse(ctx, response)
}
