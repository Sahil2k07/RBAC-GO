package handler

import (
	"net/http"
	"rbac-go/internal/config"
	interfaces "rbac-go/internal/interface"
	"rbac-go/internal/util"
	"rbac-go/internal/view"

	"github.com/labstack/echo/v4"
)

type authHandler struct {
	us interfaces.AuthService
}

func NewAuthHandler(g *echo.Group, s interfaces.AuthService) *authHandler {
	h := &authHandler{us: s}

	g.POST("/signup", h.Signup)
	g.POST("/signin", h.SignIn)

	return h
}

func (h *authHandler) Signup(c echo.Context) error {
	var req view.SignUpRequest

	if err := util.BindAndValidate(c, &req); err != nil {
		return util.HandleError(c, err)
	}

	if err := h.us.SignUp(req); err != nil {
		return util.HandleError(c, err)
	}

	return c.NoContent(http.StatusCreated)
}

func (h *authHandler) SignIn(c echo.Context) error {
	var req view.SignInRequest

	if err := util.BindAndValidate(c, &req); err != nil {
		return util.HandleError(c, err)
	}

	token, err := h.us.SignIn(req)
	if err != nil {
		return util.HandleError(c, err)
	}

	jwtConfig := config.GetJWTConfig()

	cookie := new(http.Cookie)
	cookie.Name = jwtConfig.CookieName
	cookie.Value = token
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Secure = true
	cookie.SameSite = http.SameSiteLaxMode
	cookie.MaxAge = 3600

	c.SetCookie(cookie)

	return c.NoContent(http.StatusOK)
}
