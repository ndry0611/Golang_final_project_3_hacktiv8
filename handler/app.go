package handler

import (
	"final_project_3/infrastructure/config"
	"final_project_3/infrastructure/database"

	_ "github.com/gin-gonic/gin"
)

func StartApp() {
	config.LoadAppConfig()
	db := database.GetDatabaseInstance()
	database.SeedAdmin(db)
}


