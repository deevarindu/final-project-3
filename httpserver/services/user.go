package services

import (
	"database/sql"

	"github.com/deevarindu/final-project-3/httpserver/repositories"
	"github.com/deevarindu/final-project-3/httpserver/repositories/models"
	"github.com/deevarindu/final-project-3/httpserver/views"
)

type UserSvc struct {
	repo repositories.UserRepository
}

func NewUserSvc(repo repositories.UserRepository) *UserSvc {
	return &UserSvc{
		repo: repo,
	}
}

func (u *UserSvc) GetUsers() *views.Response {
	users, err := u.repo.GetUsers()
	if err != nil {
		if err == sql.ErrNoRows {
			return views.DataNotFoundResponse(err)
		}
		return views.InternalServerErrorResponse(err)
	}

	return views.SuccessGetResponse(parseModelToGetUsers(users), "Success get all users")
}

func parseModelToGetUsers(mod *[]models.User) *[]views.GetUsers {
	var u []views.GetUsers
	for _, v := range *mod {
		u = append(u, views.GetUsers{
			ID:       *v.ID,
			FullName: v.FullName,
			Email:    v.Email,
		})
	}
	return &u
}
