package routes

import (
	"github.com/deevarindu/final-project-3/httpserver/controllers"
	"github.com/gin-gonic/gin"
)

type Router struct {
	router   *gin.Engine
	user     *controllers.UserController
	category *controllers.CategoryController
	task     *controllers.TaskController
}

func NewRouter(router *gin.Engine, user *controllers.UserController, category *controllers.CategoryController, task *controllers.TaskController) *Router {
	return &Router{
		router:   router,
		user:     user,
		category: category,
		task:     task,
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

	taskRouter := r.router.Group("/tasks")
	{
		taskRouter.GET("/", r.task.GetTasks)
	}

	r.router.Run()
}
