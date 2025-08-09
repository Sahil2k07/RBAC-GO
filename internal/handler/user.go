package handler

import (
	"net/http"
	"rbac-go/internal/authentication"
	"rbac-go/internal/enum"
	interfaces "rbac-go/internal/interface"
	"rbac-go/internal/util"
	"rbac-go/internal/view"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	us interfaces.UserService
}

func NewUserHandler(g *echo.Group, s interfaces.UserService) *userHandler {
	h := &userHandler{us: s}

	g.POST("/user", h.ListUsers, authentication.WithRole(enum.ADMIN))
	g.PUT("/user", h.UpdateUser, authentication.WithRole())
	g.GET("/user/:id", h.GetUser, authentication.WithRole())
	g.DELETE("/user/:id", h.DeleteUser, authentication.WithRole(enum.ADMIN))

	return h
}

func (h *userHandler) ListUsers(c echo.Context) error {
	var req view.ListUsers

	if err := util.BindAndValidate(c, &req); err != nil {
		return util.HandleError(c, err)
	}

	records, err := h.us.ListUsers(req)
	if err != nil {
		return util.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, records)
}

func (h *userHandler) GetUser(c echo.Context) error {
	id := c.Param("id")

	records, err := h.us.GetUser(id)
	if err != nil {
		return util.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, records)
}

func (h *userHandler) UpdateUser(c echo.Context) error {
	var req view.UserView

	if err := util.BindAndValidate(c, &req); err != nil {
		return util.HandleError(c, err)
	}

	if err := h.us.UpdateUser(req); err != nil {
		return util.HandleError(c, err)
	}

	return c.NoContent(http.StatusOK)
}

func (h *userHandler) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	if err := h.us.DeleteUser(id); err != nil {
		return util.HandleError(c, err)
	}

	return c.NoContent(http.StatusOK)
}
