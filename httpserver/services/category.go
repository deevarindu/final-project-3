package services

import (
	"database/sql"

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
		})
	}
	return &c
}
