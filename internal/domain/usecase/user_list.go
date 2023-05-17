package usecase

import (
	"context"

	"github.com/a13hander/auth-service-api/internal/domain/model"
	"github.com/a13hander/auth-service-api/internal/domain/util"
)

type listUserUseCase struct {
	repo UserRepo
	l    util.Logger
}

func NewListUserUseCase(repo UserRepo, l util.Logger) *listUserUseCase {
	return &listUserUseCase{repo: repo, l: l}
}

func (c *listUserUseCase) Run(ctx context.Context) ([]*model.User, error) {
	return c.repo.GetAll(ctx)
}
