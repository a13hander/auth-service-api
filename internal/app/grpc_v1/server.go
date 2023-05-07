package grpc_v1

import (
	"context"
	"fmt"
	"github.com/a13hander/auth-service-api/internal/domain/errs"
	"github.com/a13hander/auth-service-api/internal/domain/util"

	"github.com/a13hander/auth-service-api/internal/domain/usecase"
	desc "github.com/a13hander/auth-service-api/pkg/auth_v1"
)

type Implementation struct {
	desc.UnimplementedAuthV1Server

	createUserUseCase *usecase.CreateUserUseCase
	l                 util.Logger
}

func NewImplementation(createUserUseCase *usecase.CreateUserUseCase, l util.Logger) *Implementation {
	return &Implementation{createUserUseCase: createUserUseCase, l: l}
}

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (resp *desc.CreateResponse, err error) {
	defer func() {
		if e := recover(); e != nil {
			resp = nil
			err = errs.InternalErr

			i.l.Error(fmt.Sprintf("произошла паника при создании пользователя: %v", e))
		}
	}()

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
