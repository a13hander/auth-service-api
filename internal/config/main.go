package config

import (
	"encoding/base64"
	"log"
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	GrpcPort    string `env:"GRPC_PORT" env-required:"true"`
	HttpPort    string `env:"HTTP_PORT" env-required:"true"`
	SwaggerPort string `env:"SWAGGER_PORT" env-required:"true"`

	RefreshTokenSecretKey []byte `env:"REFRESH_TOKEN_SECRET_KEY" env-required:"true"`
	AccessTokenSecretKey  []byte `env:"ACCESS_TOKEN_SECRET_KEY" env-required:"true"`

	RefreshTokenExpirationMinutes time.Duration `env:"REFRESH_TOKEN_EXPIRATION_MINUTES" env-required:"true"`
	AccessTokenExpirationMinutes  time.Duration `env:"ACCESS_TOKEN_EXPIRATION_MINUTES" env-required:"true"`

	Db struct {
		Host     string `env:"POSTGRES_HOST" env-required:"true"`
		Port     string `env:"POSTGRES_PORT" env-required:"true"`
		User     string `env:"POSTGRES_USER" env-required:"true"`
		Password string `env:"POSTGRES_PASSWORD" env-required:"true"`
		Database string `env:"POSTGRES_DB" env-required:"true"`
	}

	RateLimit       int           `env:"RATE_LIMIT" env-required:"true"`
	RateLimitPeriod time.Duration `env:"RATE_LIMIT_PERIOD" env-required:"true"`
}

var config *Config
var onceConfig sync.Once

func GetConfig() *Config {
	onceConfig.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}

		config = &Config{}

		err = cleanenv.ReadEnv(config)
		if err != nil {
			log.Fatalln(err)
		}

		config.RefreshTokenSecretKey, err = base64.StdEncoding.DecodeString(string(config.RefreshTokenSecretKey))
		if err != nil {
			log.Fatalln(err)
		}

		config.AccessTokenSecretKey, err = base64.StdEncoding.DecodeString(string(config.AccessTokenSecretKey))
		if err != nil {
			log.Fatalln(err)
		}
	})

	return config
}
