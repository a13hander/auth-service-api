package interceptors

import (
	"context"
	"errors"
	"fmt"

	"github.com/a13hander/auth-service-api/internal/domain/errs"
	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCStatusInterface interface {
	GRPCStatus() *status.Status
}

func ErrorCodesInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (res interface{}, err error) {
	res, err = handler(ctx, req)
	if err == nil {
		return res, nil
	}

	fmt.Printf(color.RedString("error: %s\n", err.Error()))

	if de, ok := errs.AsDomainError(err); ok {
		err = handleDomainError(de)
		return
	}

	err = handleError(err)
	return
}

func handleDomainError(err errs.DomainError) error {
	grpcCode := codes.Unknown
	message := err.Error()

	switch err.Id() {
	case errs.NotFoundCode:
		grpcCode = codes.NotFound
	case errs.InvalidArgumentCode:
		grpcCode = codes.InvalidArgument
	case errs.InternalCode:
		grpcCode = codes.Internal
		message = "error"
	}

	return status.Error(grpcCode, message)
}

func handleError(err error) error {
	var se GRPCStatusInterface

	if errors.As(err, &se) {
		return se.GRPCStatus().Err()
	} else {
		if errors.Is(err, context.DeadlineExceeded) {
			err = status.Error(codes.DeadlineExceeded, err.Error())
		} else if errors.Is(err, context.Canceled) {
			err = status.Error(codes.Canceled, err.Error())
		} else {
			err = status.Error(codes.Internal, "internal error")
		}
	}

	return err
}
