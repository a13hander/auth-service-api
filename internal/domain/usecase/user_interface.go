package usecase

import (
	"context"

	"github.com/a13hander/auth-service-api/internal/domain/model"
)

type UserValidator interface {
	ValidateCreating(r *CreateUserRequest) error
}

type UserRepo interface {
	Create(ctx context.Context, u *model.User) error
	GetAll(ctx context.Context) ([]*model.User, error)
}

type CreateUserUseCase interface {
	Run(ctx context.Context, req *CreateUserRequest) (int, error)
}

type ListUserUseCase interface {
	Run(ctx context.Context) ([]*model.User, error)
}
