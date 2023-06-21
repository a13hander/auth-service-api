package access_v1

import (
	"context"

	"github.com/a13hander/auth-service-api/internal/domain/usecase"
	"github.com/a13hander/auth-service-api/internal/domain/util"
	desc "github.com/a13hander/auth-service-api/pkg/access_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Implementation struct {
	desc.UnimplementedAccessV1Server

	checkEndpoint usecase.CheckEndpoint

	l util.Logger
}

func NewImplementation(checkEndpoint usecase.CheckEndpoint, l util.Logger) *Implementation {
	return &Implementation{checkEndpoint: checkEndpoint, l: l}
}

func (i *Implementation) Check(ctx context.Context, req *desc.CheckRequest) (*emptypb.Empty, error) {
	isAllowed := i.checkEndpoint.Check(ctx, req.GetEndpoint())

	if isAllowed {
		return &emptypb.Empty{}, nil
	}

	return nil, status.Error(codes.PermissionDenied, "access denied")
}
