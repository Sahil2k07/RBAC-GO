package handler

import (
	"rbac-go/internal/repository"
	"rbac-go/internal/service"

	"github.com/labstack/echo/v4"
)

func HandleSecureEndpoints(g *echo.Group) {
	userRepo := repository.NewUserRepository()

	userService := service.NewUserService(userRepo)

	NewUserHandler(g, userService)
}
