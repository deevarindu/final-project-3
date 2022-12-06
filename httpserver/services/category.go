package services

import (
	"database/sql"

	"github.com/deevarindu/final-project-3/httpserver/controllers/params"
	"github.com/deevarindu/final-project-3/httpserver/repositories"
	"github.com/deevarindu/final-project-3/httpserver/repositories/models"
	"github.com/deevarindu/final-project-3/httpserver/views"
)

type CategorySvc struct {
	repo repositories.CategoryRepository
}

func NewCategorySvc(repo repositories.CategoryRepository) *CategorySvc {
	return &CategorySvc{
		repo: repo,
	}
}

func (c *CategorySvc) GetCategories() *views.Response {
	categories, err := c.repo.GetCategories()
	if err != nil {
		if err == sql.ErrNoRows {
			return views.DataNotFoundResponse(err)
		}
		return views.InternalServerErrorResponse(err)
	}

	return views.SuccessGetResponse(parseModelToGetCategories(categories), "Success get all categories")
}

func parseModelToGetCategories(mod *[]models.Category) *[]views.GetCategories {
	var c []views.GetCategories
	for _, v := range *mod {
		c = append(c, views.GetCategories{
			ID:   *v.ID,
			Type: v.Type,
			Tasks: func() []views.GetTasks {
				var t []views.GetTasks
				for _, v := range v.Tasks {
					t = append(t, views.GetTasks{
						ID:          *v.ID,
						Title:       v.Title,
						Description: v.Description,
						Status:      v.Status,
						UserID:      v.UserID,
						CategoryID:  v.CategoryID,
					})
				}
				return t
			}(),
		})
	}
	return &c
}

func (c *CategorySvc) CreateCategory(req *params.CategoryCreateRequest) *views.Response {
	category := parseRequestToModelCategory(req)

	id := 1
	if categories, err := c.repo.GetCategories(); err == nil {
		id = len(*categories) + 1
	}
	category.ID = &id

	err := c.repo.CreateCategory(category)

	if err != nil {
		return views.InternalServerErrorResponse(err)
	}

	return views.SuccessCreateResponse(category, "Success create category")
}

func parseRequestToModelCategory(req *params.CategoryCreateRequest) *models.Category {
	return &models.Category{
		Type: req.Type,
	}
}

func (c *CategorySvc) UpdateCategory(req *params.CategoryUpdateRequest, id int) *views.Response {
	category := parseRequestToModelUpdateCategory(req)
	category.ID = &id

	err := c.repo.UpdateCategory(category)
	if err != nil {
		return views.InternalServerErrorResponse(err)
	}

	return views.SuccessUpdateResponse(category, "Success update category")
}

func parseRequestToModelUpdateCategory(req *params.CategoryUpdateRequest) *models.Category {
	return &models.Category{
		Type: req.Type,
	}
}

func (c *CategorySvc) DeleteCategory(id int) *views.Response {
	err := c.repo.DeleteCategory(id)
	if err != nil {
		return views.InternalServerErrorResponse(err)
	}

	return views.SuccessDeleteResponse("Category has been successfully deleted")
}
