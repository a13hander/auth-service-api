package validator_test

import (
	"github.com/a13hander/auth-service-api/internal/domain/usecase"
	"github.com/a13hander/auth-service-api/internal/service/validator"
	"testing"
)

func TestPassConfirmCheck(t *testing.T) {
	sut := validator.NewUserValidator()

	request := &usecase.CreateUserRequest{
		Password:        "123",
		PasswordConfirm: "123",
	}

	err := sut.ValidateCreating(request)
	if err != nil {
		t.Fail()
	}

	request = &usecase.CreateUserRequest{
		Password:        "231",
		PasswordConfirm: "123",
	}

	err = sut.ValidateCreating(request)
	if err == nil {
		t.Fail()
	}
}
