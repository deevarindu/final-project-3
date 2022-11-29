package routes

import (
	"github.com/deevarindu/final-project-3/httpserver/controllers"
	"github.com/gin-gonic/gin"
)

type Router struct {
	router   *gin.Engine
	user     *controllers.UserController
	category *controllers.CategoryController
}

func NewRouter(router *gin.Engine, user *controllers.UserController, category *controllers.CategoryController) *Router {
	return &Router{
		router:   router,
		user:     user,
		category: category,
	}
}

func (r *Router) Start(port string) {
	userRouter := r.router.Group("/users")
	{
		userRouter.GET("/", r.user.GetUsers)
	}

	categoryRouter := r.router.Group("/categories")
	{
		categoryRouter.GET("/", r.category.GetCategories)
	}

	r.router.Run()
}
