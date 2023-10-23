package handler

import (
	"final_project_3/infrastructure/config"
	"final_project_3/infrastructure/database"
	"final_project_3/pkg/middlewares"
	"final_project_3/repository/category_repository/category_repo"
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

	categoryRepo := category_repo.NewCategoryRepo(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := NewCategoryHandler(categoryService)

	route := gin.Default()

	userRoute := route.Group("/users")
	{
		userRoute.POST("/register", userHandler.CreateUser)
		userRoute.POST("/login", userHandler.Login)

		userRoute.Use(middlewares.Authentication())
		{
			userRoute.PUT("/update-account", userHandler.UpdateUser)
		}
	}

	categoriesRoute := route.Group("/categories")
	{
		categoriesRoute.Use(middlewares.Authentication())
		{
			categoriesRoute.GET("/", categoryHandler.GetCategories)
			
			categoriesRoute.Use(middlewares.AdminAuthorization())
			categoriesRoute.POST("/", categoryHandler.CreateCategory)
		}
	}

	route.Run(":" + config.GetAppConfig().Port)
}


