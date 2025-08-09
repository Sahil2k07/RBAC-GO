package util

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	errz "rbac-go/internal/error"
)

var validate = validator.New()

func BindAndValidate(c echo.Context, req any) error {
	// Bind request
	if err := c.Bind(req); err != nil {
		return errz.NewValidation("bad request body")
	}

	// Validate request
	if err := validate.Struct(req); err != nil {
		return errz.NewValidation(err.Error())
	}

	return nil
}
