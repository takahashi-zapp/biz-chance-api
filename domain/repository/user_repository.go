package repository

import (
	"context"
	"github.com/one-connect/biz-chance-api/domain/model"
)

type UserRepository interface {
	FindAll(ctx context.Context, users []*model.User) ([]*model.User, error)
	FindOne(ctx context.Context, userID int) (*model.User, error)
	Store(ctx context.Context, user *model.User) error
}