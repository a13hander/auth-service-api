package access_v1

import (
	"context"

	"github.com/a13hander/auth-service-api/internal/domain/util"

	"github.com/a13hander/auth-service-api/internal/domain/usecase"
	desc "github.com/a13hander/auth-service-api/pkg/access_v1"
)

type Implementation struct {
	desc.UnimplementedAccessV1Server

	checkEndpoint usecase.CheckEndpoint

	l util.Logger
}

func NewImplementation(checkEndpoint usecase.CheckEndpoint, l util.Logger) *Implementation {
	return &Implementation{checkEndpoint: checkEndpoint, l: l}
}

func (i *Implementation) Check(ctx context.Context, req *desc.CheckRequest) (*desc.CheckResponse, error) {
	isAllowed := i.checkEndpoint.Check(ctx, req.GetEndpoint())

	return &desc.CheckResponse{IsAllowed: isAllowed}, nil
}
