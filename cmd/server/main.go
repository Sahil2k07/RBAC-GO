package main

import (
	"net/http"
	"rbac-go/internal/config"
	"rbac-go/internal/database"
	"rbac-go/internal/handler"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	configs := config.LoadConfig()
	database.Connect()

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     configs.Origins,
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))

	public := e.Group("public")
	handler.HandlePublicEndpoints(public)

	secure := e.Group("api/v1")
	secure.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(configs.JWT.Secret),
		TokenLookup: "cookie:" + configs.JWT.CookieName,
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid or expired token"})
		},
	}))
	handler.HandleSecureEndpoints(secure)

	e.Logger.Fatal(e.Start(":5000"))
}
