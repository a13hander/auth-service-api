package usecase

import (
	"context"
	"fmt"
	"github.com/a13hander/auth-service-api/internal/domain/errs"
	"github.com/a13hander/auth-service-api/internal/domain/model"
	"github.com/a13hander/auth-service-api/internal/domain/util"
	"time"
)

type UserRepo interface {
	Create(ctx context.Context, u *model.User) error
}

type UserValidator interface {
	ValidateCreating(r *CreateUserRequest) error
}

type CreateUserRequest struct {
	Email           string
	Username        string
	Password        string
	PasswordConfirm string
	Role            string
}

func (r *CreateUserRequest) String() string {
	return fmt.Sprintf("%v", *r)
}

type CreateUserUseCase struct {
	validator UserValidator
	repo      UserRepo
	l         util.Logger
}

func NewCreateUserUseCase(validator UserValidator, repo UserRepo, l util.Logger) *CreateUserUseCase {
	return &CreateUserUseCase{validator: validator, repo: repo, l: l}
}

func (c *CreateUserUseCase) Create(ctx context.Context, req *CreateUserRequest) (int, error) {
	err := c.validator.ValidateCreating(req)
	if err != nil {
		return 0, errs.NewInvalidArgumentError(err.Error())
	}

	u := model.User{}
	fillAttrs(&u, req)

	err = c.repo.Create(ctx, &u)
	if err != nil {
		c.l.ErrorWithCtx("не удалось создать пользователя", map[string]any{
			"err":     err.Error(),
			"payload": req.String(),
		})
		return 0, errs.InternalErr
	}

	return u.Id, nil
}

func fillAttrs(u *model.User, req *CreateUserRequest) {
	now := time.Now()

	u.Email = req.Email
	u.Username = req.Username
	u.Password = req.Password
	u.Role = req.Role
	u.CreatedAt = now
	u.UpdatedAt = now
}
