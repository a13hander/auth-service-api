package interceptors

import (
	"context"

	"google.golang.org/grpc"
)

type validator interface {
	Validate() error
}

func ValidateInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	if v, ok := req.(validator); ok {
		err := v.Validate()
		if err != nil {
			return nil, err
		}
	}

	return handler(ctx, req)
}
