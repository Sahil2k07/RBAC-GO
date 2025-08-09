package interfaces

import (
	"rbac-go/internal/model"
	"rbac-go/internal/view"
)

type (
	AuthRepository interface {
		CheckUserExist(email string) (bool, error)

		GetUser(email string) (model.User, error)

		AddUser(req view.SignUpRequest) error
	}

	AuthService interface {
		SignUp(req view.SignUpRequest) error

		SignIn(req view.SignInRequest) (string, error)
	}
)
