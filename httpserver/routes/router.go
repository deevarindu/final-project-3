package routes

import (
	"github.com/deevarindu/final-project-3/httpserver/controllers"
	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
	user   *controllers.UserController
}

func NewRouter(router *gin.Engine, user *controllers.UserController) *Router {
	return &Router{
		router: router,
		user:   user,
	}
}

func (r *Router) Start(port string) {
	userRouter := r.router.Group("/users")
	{
		userRouter.GET("/", r.user.GetUsers)
	}

	r.router.Run()
}
