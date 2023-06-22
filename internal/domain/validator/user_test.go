package validator_test

import (
	"context"
	"testing"

	"github.com/a13hander/auth-service-api/internal/domain/usecase"
	"github.com/a13hander/auth-service-api/internal/domain/validator"
)

func TestPassConfirmCheck(t *testing.T) {
	sut := validator.NewUserValidator()

	request := &usecase.CreateUserRequest{
		Password:        "1234567890",
		PasswordConfirm: "1234567890",
	}

	err := sut.ValidateCreating(context.Background(), request)
	if err != nil {
		t.Fail()
	}

	request = &usecase.CreateUserRequest{
		Password:        "1234567890",
		PasswordConfirm: "0234567891",
	}

	err = sut.ValidateCreating(context.Background(), request)
	if err == nil {
		t.Fail()
	}
}
