package usecase

import (
	"context"
	"github.com/one-connect/biz-chance-api/domain/model"
	"github.com/one-connect/biz-chance-api/domain/repository"
	"github.com/one-connect/biz-chance-api/interfaces/presenter"
)

type userUseCase struct {
	repository.UserRepository
	presenter.UserPresenter
}

type UserUseCase interface {
	Create(ctx context.Context) (*model.Users, error)
	Get(ctx context.Context, id int) (*model.User, error)
}

func NewUserUseCase(repo repository.UserRepository, pre presenter.UserPresenter) userUseCase {
	return userUseCase{repo, pre}
}

func (userUseCase *userUseCase) Create(ctx context.Context, id int) error {
	return userUseCase.UserRepository.Store(u)
}

func (userUseCase *userUseCase) Get(ctx context.Context) ([]*model.User, error) {
	us, err := userUseCase.UserRepository.FindAll(u)
	if err != nil {
		return nil, err
	}
	return userUseCase.UserPresenter.ResponseUsers(us), nil
}