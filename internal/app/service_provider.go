package app

import (
	"context"
	"fmt"
	"log"

	grpcV1 "github.com/a13hander/auth-service-api/internal/app/grpc_v1"
	"github.com/a13hander/auth-service-api/internal/config"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/a13hander/auth-service-api/internal/domain/usecase"
	"github.com/a13hander/auth-service-api/internal/domain/util"
	"github.com/a13hander/auth-service-api/internal/domain/validator"
	"github.com/a13hander/auth-service-api/internal/service/database"
	"github.com/a13hander/auth-service-api/internal/service/logger"
)

type serviceProvider struct {
	logger           util.Logger
	pool             *pgxpool.Pool
	grpcV1ServerImpl *grpcV1.Implementation

	repo struct {
		userRepo usecase.UserRepo
	}

	validator struct {
		userValidator usecase.UserValidator
	}

	useCase struct {
		createUserUseCase *usecase.CreateUserUseCase
	}
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (c *serviceProvider) GetLogger(_ context.Context) util.Logger {
	if c.logger == nil {
		c.logger = logger.NewLogger()
	}

	return c.logger
}

func (c *serviceProvider) GetPgxPool(ctx context.Context) *pgxpool.Pool {
	if c.pool == nil {
		dbConf := config.GetConfig().Db

		conf, _ := pgxpool.ParseConfig(fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
			dbConf.User,
			dbConf.Password,
			dbConf.Host,
			dbConf.Port,
			dbConf.Database),
		)

		pool, err := pgxpool.ConnectConfig(ctx, conf)
		if err != nil {
			log.Fatalln(err)
		}

		c.pool = pool
	}

	return c.pool
}

func (c *serviceProvider) GetUserRepo(ctx context.Context) usecase.UserRepo {
	if c.repo.userRepo == nil {
		c.repo.userRepo = database.NewUserRepo(c.GetPgxPool(ctx))
	}

	return c.repo.userRepo
}

func (c *serviceProvider) GetUserValidator(_ context.Context) usecase.UserValidator {
	if c.validator.userValidator == nil {
		c.validator.userValidator = validator.NewUserValidator()
	}

	return c.validator.userValidator
}

func (c *serviceProvider) GetCreateUserUseCase(ctx context.Context) *usecase.CreateUserUseCase {
	if c.useCase.createUserUseCase == nil {
		c.useCase.createUserUseCase = usecase.NewCreateUserUseCase(c.GetUserValidator(ctx), c.GetUserRepo(ctx), c.GetLogger(ctx))
	}

	return c.useCase.createUserUseCase
}

func (c *serviceProvider) GetGrpcV1ServerImpl(ctx context.Context) *grpcV1.Implementation {
	if c.grpcV1ServerImpl == nil {
		c.grpcV1ServerImpl = grpcV1.NewImplementation(c.GetCreateUserUseCase(ctx), c.GetLogger(ctx))
	}

	return c.grpcV1ServerImpl
}
