package access_v1

import (
	"context"
	"strings"

	"github.com/a13hander/auth-service-api/internal/domain/usecase"
	"github.com/a13hander/auth-service-api/internal/domain/util"
	desc "github.com/a13hander/auth-service-api/pkg/access_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

const authPrefix = "Bearer "

type Implementation struct {
	desc.UnimplementedAccessV1Server

	checkEndpoint usecase.CheckEndpoint

	l util.Logger
}

func NewImplementation(checkEndpoint usecase.CheckEndpoint, l util.Logger) *Implementation {
	return &Implementation{checkEndpoint: checkEndpoint, l: l}
}

func (i *Implementation) Check(ctx context.Context, req *desc.CheckRequest) (*emptypb.Empty, error) {
	accessDenied := status.Error(codes.PermissionDenied, "access denied")

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, accessDenied
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return nil, accessDenied
	}

	if !strings.HasPrefix(authHeader[0], authPrefix) {
		return nil, accessDenied
	}

	token := strings.TrimPrefix(authHeader[0], authPrefix)
	isAllowed := i.checkEndpoint.Check(ctx, token, req.GetEndpoint())

	if isAllowed {
		return &emptypb.Empty{}, nil
	}

	return nil, accessDenied
}
