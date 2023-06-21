package validator

import (
	"context"

	"github.com/a13hander/auth-service-api/internal/domain/errs"
	"github.com/a13hander/auth-service-api/internal/domain/usecase"
)

type UserValidator struct {
}

func NewUserValidator() *UserValidator {
	return &UserValidator{}
}

func (v *UserValidator) ValidateCreating(ctx context.Context, r *usecase.CreateUserRequest) error {
	return errs.Validate(ctx,
		ValidatePassword(r.Password, r.PasswordConfirm),
		ValidatePasswordLength(r.Password),
	)
}

func ValidatePassword(password, passwordConfirm string) errs.Condition {
	return func(ctx context.Context) error {
		if password != passwordConfirm {
			return errs.NewInvalidArgumentError("пароль и подтверждение не совпадают")
		}

		return nil
	}
}

func ValidatePasswordLength(password string) errs.Condition {
	return func(ctx context.Context) error {
		passLen := len(password)
		if passLen < 8 || passLen > 32 {
			return errs.NewInvalidArgumentError("пароль должен иметь длинну от 8 до 32 символов")
		}

		return nil
	}
}
