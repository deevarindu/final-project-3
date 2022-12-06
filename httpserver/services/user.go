package services

import (
	"database/sql"
	"strings"

	"github.com/deevarindu/final-project-3/httpserver/controllers/params"
	"github.com/deevarindu/final-project-3/httpserver/repositories"
	"github.com/deevarindu/final-project-3/httpserver/repositories/models"
	"github.com/deevarindu/final-project-3/httpserver/views"
	"golang.org/x/crypto/bcrypt"
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

func (u *UserSvc) Register(req *params.UserRegisterRequest) *views.Response {
	user := parseRequestToModelRegister(req)
	user.Password = hashedPassword(user.Password)

	id := 1
	if users, err := u.repo.GetUsers(); err == nil {
		id = len(*users) + 1
	}
	user.ID = &id

	err := u.repo.Register(user)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return views.DataConflictResponse(err)
		}
		return views.InternalServerErrorResponse(err)
	}

	return views.SuccessCreateResponse(user, "Success register user")
}

func hashedPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

func parseRequestToModelRegister(req *params.UserRegisterRequest) *models.User {
	return &models.User{
		FullName: req.FullName,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}
}

func (u *UserSvc) FindUserByEmail(email string) *models.User {
	users, _ := u.repo.GetUsers()
	for _, user := range *users {
		if strings.EqualFold(user.Email, email) {
			return &user
		}
	}
	return nil
}

func (u *UserSvc) FindUserById(id *int) *models.User {
	users, _ := u.repo.GetUsers()
	for _, user := range *users {
		if *user.ID == *id {
			return &user
		}
	}
	return nil
}

func (u *UserSvc) UpdateUser(req *params.UserUpdateRequest, id *int) *views.Response {
	user := u.FindUserById(id)
	if user == nil {
		return views.DataNotFoundResponse(nil)
	}

	user.FullName = req.FullName
	user.Email = req.Email

	err := u.repo.UpdateUser(user)
	if err != nil {
		return views.InternalServerErrorResponse(err)
	}

	return views.SuccessUpdateResponse(user, "Success update user")
}

func (u *UserSvc) DeleteUser(id *int) *views.Response {
	user := u.FindUserById(id)
	if user == nil {
		return views.DataNotFoundResponse(nil)
	}

	err := u.repo.DeleteUser(user)
	if err != nil {
		return views.InternalServerErrorResponse(err)
	}

	return views.SuccessDeleteResponse("Your account has been success deleted")
}
