package container

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	GrpcPort string `env:"GRPC_PORT" env-required:"true"`

	Db struct {
		Host     string `env:"POSTGRES_HOST" env-required:"true"`
		Port     string `env:"POSTGRES_PORT" env-required:"true"`
		User     string `env:"POSTGRES_USER" env-required:"true"`
		Password string `env:"POSTGRES_PASSWORD" env-required:"true"`
		Database string `env:"POSTGRES_DB" env-required:"true"`
	}
}

var config *Config
var onceConfig sync.Once

func GetConfig() *Config {
	onceConfig.Do(func() {
		config = &Config{}

		err := cleanenv.ReadEnv(config)
		if err != nil {
			log.Fatalln(err)
		}
	})

	return config
}
