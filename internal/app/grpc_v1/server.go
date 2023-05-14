package grpc_v1

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/a13hander/auth-service-api/internal/domain/errs"
	"github.com/a13hander/auth-service-api/internal/domain/util"

	"github.com/a13hander/auth-service-api/internal/domain/usecase"
	desc "github.com/a13hander/auth-service-api/pkg/auth_v1"
)

type Implementation struct {
	desc.UnimplementedAuthV1Server

	createUserUseCase usecase.CreateUserUseCase
	listUserUseCase   usecase.ListUserUseCase
	l                 util.Logger
}

func NewImplementation(createUserUseCase usecase.CreateUserUseCase, listUserUseCase usecase.ListUserUseCase, l util.Logger) *Implementation {
	return &Implementation{createUserUseCase: createUserUseCase, listUserUseCase: listUserUseCase, l: l}
}

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (resp *desc.CreateResponse, err error) {
	// TODO перенести в мидлваре
	defer func() {
		if e := recover(); e != nil {
			resp = nil
			err = errs.InternalErr

			i.l.Error(fmt.Sprintf("произошла паника при создании пользователя: %v", e))
		}
	}()

	userInfo := req.GetUser()
	createReq := usecase.CreateUserRequest{
		Email:           userInfo.GetEmail(),
		Username:        userInfo.GetUsername(),
		Password:        req.GetPassword(),
		PasswordConfirm: req.GetPasswordConfirm(),
		Role:            int(userInfo.GetRole()),
	}

	id, err := i.createUserUseCase.Run(ctx, &createReq)
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{Id: uint32(id)}, nil
}

func (i *Implementation) List(ctx context.Context, _ *emptypb.Empty) (resp *desc.ListResponse, err error) {
	// TODO перенести в мидлваре
	defer func() {
		if e := recover(); e != nil {
			resp = nil
			err = errs.InternalErr

			i.l.Error(fmt.Sprintf("произошла паника при получение пользователей: %v", e))
		}
	}()

	users, err := i.listUserUseCase.Run(ctx)
	if err != nil {
		return nil, err
	}

	descUsers := make([]*desc.User, 0, len(users))
	for _, u := range users {
		info := &desc.UserInfo{
			Email:    u.Email,
			Username: u.Username,
			Role:     desc.Role(u.Role),
		}
		descUsers = append(descUsers, &desc.User{
			Id:        uint32(u.Id),
			Info:      info,
			CreatedAt: timestamppb.New(u.CreatedAt),
		})
	}

	return &desc.ListResponse{User: descUsers}, nil
}
