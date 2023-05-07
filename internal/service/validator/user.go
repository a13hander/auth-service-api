package validator

import (
	"errors"
	"github.com/a13hander/auth-service-api/internal/domain/usecase"
)

type UserValidator struct {
}

func NewUserValidator() *UserValidator {
	return &UserValidator{}
}

func (v *UserValidator) ValidateCreating(r *usecase.CreateUserRequest) error {
	if r.Password != r.PasswordConfirm {
		return errors.New("пароль и подтверждение не совпадают")
	}

	return nil
}
