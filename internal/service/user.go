package service

import (
	interfaces "rbac-go/internal/interface"
	"rbac-go/internal/model"
	"rbac-go/internal/util"
	"rbac-go/internal/view"
)

type userService struct {
	repo interfaces.UserRepository
}

func NewUserService(r interfaces.UserRepository) interfaces.UserService {
	return &userService{repo: r}
}

func (s *userService) ListUsers(req view.ListUsers) (view.ListResponse, error) {
	users, count, err := s.repo.ListUsers(req)
	if err != nil {
		return view.ListResponse{}, err
	}

	records := util.Map(users, func(m model.User) view.UserView {
		return view.NewUserResponse(m)
	})

	resp := view.ListResponse{
		TotalRecords: count,
		Data:         records,
	}

	return resp, nil
}

func (s *userService) GetUser(id string) (view.UserView, error) {
	uid, err := util.StringToUint(id)
	if err != nil {
		return view.UserView{}, err
	}

	user, err := s.repo.GetUser(uid)
	if err != nil {
		return view.UserView{}, err
	}

	return view.NewUserResponse(user), nil

}

func (s *userService) UpdateUser(req view.UserView) error {
	if err := s.repo.UpdateUser(req); err != nil {
		return err
	}

	return nil
}

func (s *userService) DeleteUser(id string) error {
	uid, err := util.StringToUint(id)
	if err != nil {
		return err
	}

	if err := s.repo.DeleteUser(uid); err != nil {
		return err
	}

	return nil
}
