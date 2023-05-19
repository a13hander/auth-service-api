package usecase

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/a13hander/auth-service-api/internal/domain/errs"
	"github.com/a13hander/auth-service-api/internal/domain/model"
	"github.com/a13hander/auth-service-api/internal/domain/util"
)

type CreateUserRequest struct {
	Email           string
	Username        string
	Password        string
	PasswordConfirm string
	Role            int
	// TODO refactor
	Engineer *struct {
		Level    int64
		Company  string
		Language string
	}
	Manager *struct {
		Level      int64
		Company    string
		Experience int64
	}
}

func (r *CreateUserRequest) String() string {
	return fmt.Sprintf("%v", *r)
}

type createUserUseCase struct {
	validator UserValidator
	repo      UserRepo
	l         util.Logger
}

func NewCreateUserUseCase(validator UserValidator, repo UserRepo, l util.Logger) *createUserUseCase {
	return &createUserUseCase{validator: validator, repo: repo, l: l}
}

func (c *createUserUseCase) Run(ctx context.Context, req *CreateUserRequest) (int, error) {
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
	u.Email = req.Email
	u.Username = req.Username
	u.Password = req.Password
	u.Role = req.Role

	if req.Engineer != nil {
		data, _ := json.Marshal(req.Engineer)
		u.Specialisation = data
	}

	if req.Manager != nil {
		data, _ := json.Marshal(req.Manager)
		u.Specialisation = data
	}
}
