package handler

import (
	"final_project_3/infrastructure/config"
	"final_project_3/infrastructure/database"
	"final_project_3/repository/user_repository/user_repo"
	"final_project_3/service"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	config.LoadAppConfig()
	db := database.GetDatabaseInstance()
	database.SeedAdmin(db)

	//Dependency Injection
	userRepo := user_repo.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	route := gin.Default()

	userRoute := route.Group("/users")
	{
		userRoute.POST("/register", userHandler.CreateUser)
		userRoute.POST("/login", userHandler.Login)
	}

	route.Run(":" + config.GetAppConfig().Port)
}


