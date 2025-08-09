package handler

import (
	"rbac-go/internal/repository"
	"rbac-go/internal/service"

	"github.com/labstack/echo/v4"
)

func HandlePublicEndpoints(g *echo.Group) {
	authRepo := repository.NewAuthRepository()

	authService := service.NewAuthService(authRepo)

	NewAuthHandler(g, authService)
}
