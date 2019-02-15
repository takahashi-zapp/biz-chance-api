package service

import (
	"context"
	"github.com/one-connect/biz-chance-api/domain/model"
	"github.com/one-connect/biz-chance-api/domain/repository"
	"github.com/one-connect/biz-chance-api/interfaces/presenter"
)

type userService struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
}

type UserService interface {
	Create(u *model.User) error
	Get(u []*model.User) ([]*model.User, error)
}

func NewUserService(repo repository.UserRepository, pre presenter.UserPresenter) UserService {
	return &userService{repo, pre}
}

func (userService *userService) Create(ctx context.Context, u *model.User) error {
	return userService.UserRepository.Store(ctx, u)
}

func (userService *userService) Get(ctx context.Context, u []*model.User) ([]*model.User, error) {
	us, err := userService.UserRepository.FindAll(ctx, u)
	if err != nil {
		return nil, err
	}
	return userService.UserPresenter.ResponseUsers(us), nil
}