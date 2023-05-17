package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

func createPgxPool(ctx context.Context, dbConf DbConfig) *pgxpool.Pool {
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

	err = pool.Ping(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return pool
}
