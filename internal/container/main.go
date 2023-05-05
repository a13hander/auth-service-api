package container

import (
	"context"
	"fmt"
	"log"

	"github.com/a13hander/auth-service-api/internal/app/grpc_v1"
	"github.com/a13hander/auth-service-api/internal/domain/usecase"
	"github.com/a13hander/auth-service-api/internal/domain/util"
	"github.com/a13hander/auth-service-api/internal/service/database"
	"github.com/a13hander/auth-service-api/internal/service/logger"
	"github.com/a13hander/auth-service-api/internal/service/validator"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

type Container struct {
	pool   *pgxpool.Pool
	Logger util.Logger

	Repo struct {
		UserRepo usecase.UserRepo
	}

	UseCase struct {
		CreateUserUseCase *usecase.CreateUserUseCase
	}

	Grpc struct {
		V1 *grpc_v1.Implementation
	}
}

func Build() *Container {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()
	appConfig := GetConfig()
	cont := Container{}

	config, _ := pgxpool.ParseConfig(fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		appConfig.Db.User,
		appConfig.Db.Password,
		appConfig.Db.Host,
		appConfig.Db.Port,
		appConfig.Db.Database),
	)

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		log.Fatalln(err)
	}

	cont.pool = pool
	cont.Logger = logger.NewLogger()

	cont.Repo.UserRepo = database.NewUserRepo(pool)

	cont.UseCase.CreateUserUseCase = usecase.NewCreateUserUseCase(validator.NewUserValidator(), cont.Repo.UserRepo, cont.Logger)

	cont.Grpc.V1 = grpc_v1.NewImplementation(cont.UseCase.CreateUserUseCase, cont.Logger)

	return &cont
}
