package controller

import (
	"github.com/one-connect/biz-chance-api/application/usecase"
	"github.com/one-connect/biz-chance-api/domain/model"
	"github.com/one-connect/biz-chance-api/domain/service"
)

type userController struct {
	userUseCase usecase.UserUseCase
}

type UserController interface {
	CreateUser(user *model.User) error
	GetUsers() ([]*model.User, error)
}

func NewUserController(us service.UserService) UserController {
	return &userController{us}
}

func (userController *userController) CreateUser(user *model.User) error {
	// 内側の処理のための技術的な関心事を記載
	return userController.userService.Create(user)
}

func (userController *userController) GetUsers() ([]*model.User, error) {
	u := []*model.User{}
	us, err := userController.userService.Get(u)
	if err != nil {
		return nil, err
	}
	return us, nil
}