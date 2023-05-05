package grpc_v1

import (
	"context"

	"github.com/a13hander/auth-service-api/internal/domain/usecase"
	desc "github.com/a13hander/auth-service-api/pkg/auth_v1"
)

type Implementation struct {
	desc.UnimplementedAuthV1Server

	createUserUseCase *usecase.CreateUserUseCase
}

func NewImplementation(createUserUseCase *usecase.CreateUserUseCase) *Implementation {
	return &Implementation{createUserUseCase: createUserUseCase}
}

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	createReq := usecase.CreateUserRequest{
		Email:           req.GetEmail(),
		Username:        req.GetUsername(),
		Password:        req.GetPassword(),
		PasswordConfirm: req.GetPasswordConfirm(),
		Role:            int(req.GetRole()),
	}

	id, err := i.createUserUseCase.Run(ctx, &createReq)
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{Id: uint32(id)}, nil
}
