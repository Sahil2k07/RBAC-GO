package main

import (
	"rbac-go/internal/config"
	"rbac-go/internal/database"
	"rbac-go/internal/model"

	"github.com/charmbracelet/log"
)

func main() {
	config.LoadConfig()
	database.Connect()

	// Migration - 1
	models := []any{
		&model.User{}, &model.Profile{},
	}

	err := database.DB.AutoMigrate(models...)
	if err != nil {
		log.Errorf("Migration failed: %v", err)
		return
	}

	log.Infof("Migration Completed Successfully!")
}
