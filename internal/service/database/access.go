package database

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"github.com/a13hander/auth-service-api/internal/domain/model"
	"github.com/a13hander/auth-service-api/internal/domain/usecase"
)

const accessTableName = "access"

type accessRepo struct {
	dbClient Client
}

func NewAccessRepo(dbClient Client) usecase.AccessRepo {
	return &accessRepo{
		dbClient: dbClient,
	}
}

func (r *accessRepo) Get(ctx context.Context, endpoint string) (*model.Access, error) {
	sql, args, err := sq.Select("id", "endpoint", "role", "created_at", "updated_at").
		From(accessTableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"endpoint": endpoint}).
		Limit(1).
		ToSql()

	if err != nil {
		return nil, err
	}

	q := Query{
		Name:     "AccessRepo.Get",
		QueryRaw: sql,
	}

	access := &model.Access{}
	err = r.dbClient.Get(ctx, access, q, args...)
	if err != nil {
		return nil, err
	}

	return access, nil
}
