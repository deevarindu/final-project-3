package main

import (
	"github.com/deevarindu/final-project-3/config"
	"github.com/deevarindu/final-project-3/httpserver/controllers"
	"github.com/deevarindu/final-project-3/httpserver/repositories/gorm"
	"github.com/deevarindu/final-project-3/httpserver/routes"
	"github.com/deevarindu/final-project-3/httpserver/services"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.CreateConnection()
	if err != nil {
		panic(err)
	}

	userRepository := gorm.NewUserRepository(db)
	userSvc := services.NewUserSvc(userRepository)
	userHandler := controllers.NewUserController(userSvc)

	categoryRepository := gorm.NewCategoryRepository(db)
	categorySvc := services.NewCategorySvc(categoryRepository)
	categoryHandler := controllers.NewCategoryController(categorySvc)

	router := gin.Default()
	app := routes.NewRouter(router, userHandler, categoryHandler)

	app.Start(":5000")
}
