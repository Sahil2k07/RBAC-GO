package main

import (
	"rabc-go/internal/database"

	"github.com/charmbracelet/log"
)

func main() {
	database.Connect()

	err := database.DB.AutoMigrate()
	if err != nil {
		log.Errorf("Migration failed: %v", err)
	}

	log.Infof("Migration Completed Successfully!")
}
