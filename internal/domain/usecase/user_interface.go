package usecase

import (
	"context"

	"github.com/a13hander/auth-service-api/internal/domain/model"
)

type UserValidator interface {
	ValidateCreating(ctx context.Context, r *CreateUserRequest) error
}

type UserRepo interface {
	Get(ctx context.Context, username string) (*model.User, error)
	Create(ctx context.Context, u *model.User) error
	GetAll(ctx context.Context) ([]*model.User, error)
}

type CreateUserUseCase interface {
	Run(ctx context.Context, req *CreateUserRequest) (int, error)
}

type ListUserUseCase interface {
	Run(ctx context.Context) ([]*model.User, error)
}

type RefreshTokenGenerator interface {
	Generate(ctx context.Context, username string, password string) (string, error)
}

type AccessTokenGenerator interface {
	Generate(ctx context.Context, refreshToken string) (string, error)
}
