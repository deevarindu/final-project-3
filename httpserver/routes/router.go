package routes

import (
	"github.com/deevarindu/final-project-3/httpserver/controllers"
	"github.com/deevarindu/final-project-3/httpserver/middleware"
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
		userRouter.POST("/register", r.user.Register)
		userRouter.POST("/login", r.user.Login)
		userRouter.Use(middleware.Authentication())
		userRouter.PUT("/update-account", r.user.UpdateUser)
		userRouter.DELETE("/delete-account", r.user.DeleteUser)
	}

	categoryRouter := r.router.Group("/categories")
	{
		categoryRouter.GET("/", r.category.GetCategories)
		categoryRouter.Use(middleware.Authentication(), middleware.Authorization)
		categoryRouter.POST("/", r.category.CreateCategory)
		categoryRouter.PATCH("/:categoryId", r.category.UpdateCategory)
		categoryRouter.DELETE("/:categoryId", r.category.DeleteCategory)
	}

	taskRouter := r.router.Group("/tasks")
	{
		taskRouter.GET("/", r.task.GetTasks)
		taskRouter.Use(middleware.Authentication())
		taskRouter.POST("/", r.task.CreateTask)
		taskRouter.PUT("/:taskId", r.task.UpdateTitleDesc)
		taskRouter.PATCH("/update-status/:taskId", r.task.UpdateStatus)
		taskRouter.PATCH("/update-category/:taskId", r.task.UpdateCategory)
		taskRouter.DELETE("/:taskId", r.task.DeleteTask)
	}

	r.router.Run(port)
}
