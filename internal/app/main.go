package app

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/a13hander/auth-service-api/internal/config"
	desc "github.com/a13hander/auth-service-api/pkg/auth_v1"
)

type App struct {
	serviceProvider *serviceProvider
	config          *config.Config
	grpcV1Server    *grpc.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run(ctx context.Context) error {
	return a.runGrpcV1Server(ctx)
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGrpcV1Server,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	a.config = config.GetConfig()
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initGrpcV1Server(ctx context.Context) error {
	a.grpcV1Server = grpc.NewServer()
	reflection.Register(a.grpcV1Server)
	desc.RegisterAuthV1Server(a.grpcV1Server, a.serviceProvider.GetGrpcV1ServerImpl(ctx))

	return nil
}

func (a *App) runGrpcV1Server(_ context.Context) error {
	listener, err := net.Listen("tcp", a.config.GrpcPort)
	if err != nil {
		log.Fatalln(err)
	}

	err = a.grpcV1Server.Serve(listener)
	if err != nil {
		return err
	}

	return nil
}
