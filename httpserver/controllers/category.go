package controllers

import (
	"net/http"
	"strconv"

	"github.com/deevarindu/final-project-3/httpserver/controllers/params"
	"github.com/deevarindu/final-project-3/httpserver/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	response := c.svc.GetCategories()
	WriteJsonResponse(ctx, response)
}

func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	var req params.CategoryCreateRequest
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

	response := c.svc.CreateCategory(&req)
	WriteJsonResponse(ctx, response)
}

func (c *CategoryController) UpdateCategory(ctx *gin.Context) {
	var req params.CategoryUpdateRequest
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

	categoryID, _ := strconv.Atoi(ctx.Param("categoryId"))
	response := c.svc.UpdateCategory(&req, categoryID)
	WriteJsonResponse(ctx, response)
}

func (c *CategoryController) DeleteCategory(ctx *gin.Context) {
	categoryID, _ := strconv.Atoi(ctx.Param("categoryId"))
	response := c.svc.DeleteCategory(categoryID)
	WriteJsonResponse(ctx, response)
}
