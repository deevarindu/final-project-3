package controllers

import (
	"github.com/deevarindu/final-project-3/httpserver/services"
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	svc *services.CategorySvc
}

func NewCategoryController(svc *services.CategorySvc) *CategoryController {
	return &CategoryController{
		svc: svc,
	}
}

func (c *CategoryController) GetCategories(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	response := c.svc.GetCategories()
	WriteJsonResponse(ctx, response)
}
