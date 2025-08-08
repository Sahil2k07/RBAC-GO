package main

import (
	"rabc-go/internal/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.Connect()

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.Logger.Fatal(e.Start(":5000"))
}
