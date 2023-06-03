package app

import (
	"context"

	grpcV1 "github.com/a13hander/auth-service-api/internal/app/grpc_v1"
	"github.com/a13hander/auth-service-api/internal/config"

	"github.com/a13hander/auth-service-api/internal/domain/usecase"
	"github.com/a13hander/auth-service-api/internal/domain/util"
	"github.com/a13hander/auth-service-api/internal/domain/validator"
	"github.com/a13hander/auth-service-api/internal/service/database"
	"github.com/a13hander/auth-service-api/internal/service/logger"
)

type serviceProvider struct {
	logger           util.Logger
	dbClient         database.Client
	grpcV1ServerImpl *grpcV1.Implementation

	repo struct {
		userRepo usecase.UserRepo
	}

	validator struct {
		userValidator usecase.UserValidator
	}

	useCase struct {
		createUserUseCase usecase.CreateUserUseCase
		userListUseCase   usecase.ListUserUseCase

		refreshTokenGenerator usecase.RefreshTokenGenerator
		accessTokenGenerator  usecase.AccessTokenGenerator
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

func (c serviceProvider) GetDbClient(ctx context.Context) database.Client {
	if c.dbClient == nil {
		dbConf := config.GetConfig().Db

		c.dbClient = database.NewClient(ctx, database.DbConfig{
			Host:     dbConf.Host,
			Port:     dbConf.Port,
			User:     dbConf.User,
			Password: dbConf.Password,
			Database: dbConf.Database,
		})

		closer.add(c.dbClient.Close)
	}

	return c.dbClient
}

func (c *serviceProvider) GetUserRepo(ctx context.Context) usecase.UserRepo {
	if c.repo.userRepo == nil {
		c.repo.userRepo = database.NewUserRepo(c.GetDbClient(ctx))
	}

	return c.repo.userRepo
}

func (c *serviceProvider) GetUserValidator(_ context.Context) usecase.UserValidator {
	if c.validator.userValidator == nil {
		c.validator.userValidator = validator.NewUserValidator()
	}

	return c.validator.userValidator
}

func (c *serviceProvider) GetCreateUserUseCase(ctx context.Context) usecase.CreateUserUseCase {
	if c.useCase.createUserUseCase == nil {
		c.useCase.createUserUseCase = usecase.NewCreateUserUseCase(c.GetUserValidator(ctx), c.GetUserRepo(ctx), c.GetLogger(ctx))
	}

	return c.useCase.createUserUseCase
}

func (c *serviceProvider) GetListUserUseCase(ctx context.Context) usecase.ListUserUseCase {
	if c.useCase.userListUseCase == nil {
		c.useCase.userListUseCase = usecase.NewListUserUseCase(c.GetUserRepo(ctx), c.GetLogger(ctx))
	}

	return c.useCase.userListUseCase
}

func (c *serviceProvider) GetRefreshTokenGenerator(ctx context.Context) usecase.RefreshTokenGenerator {
	if c.useCase.refreshTokenGenerator == nil {
		conf := config.GetConfig()
		c.useCase.refreshTokenGenerator = usecase.NewRefreshTokenGenerator(c.GetUserRepo(ctx), conf.RefreshTokenSecretKey, conf.RefreshTokenExpirationMinutes)
	}

	return c.useCase.refreshTokenGenerator
}

func (c *serviceProvider) GetAccessTokenGenerator(ctx context.Context) usecase.AccessTokenGenerator {
	if c.useCase.accessTokenGenerator == nil {
		conf := config.GetConfig()
		c.useCase.accessTokenGenerator = usecase.NewAccessTokenGenerator(c.GetUserRepo(ctx), conf.RefreshTokenSecretKey, conf.AccessTokenSecretKey, conf.AccessTokenExpirationMinutes)
	}

	return c.useCase.accessTokenGenerator
}

func (c *serviceProvider) GetGrpcV1ServerImpl(ctx context.Context) *grpcV1.Implementation {
	if c.grpcV1ServerImpl == nil {
		c.grpcV1ServerImpl = grpcV1.NewImplementation(
			c.GetCreateUserUseCase(ctx),
			c.GetListUserUseCase(ctx),
			c.GetRefreshTokenGenerator(ctx),
			c.GetAccessTokenGenerator(ctx),
			c.GetLogger(ctx),
		)
	}

	return c.grpcV1ServerImpl
}
