package database

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/a13hander/auth-service-api/internal/domain/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

const tableName = "users"

type UserRepo struct {
	pool *pgxpool.Pool
}

func NewUserRepo(pool *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		pool: pool,
	}
}

func (r *UserRepo) Create(ctx context.Context, u *model.User) error {
	sql, v, err := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns("email", "username", "password", "role", "created_at", "updated_at").
		Values(u.Email, u.Username, u.Password, u.Role, u.CreatedAt, u.UpdatedAt).
		Suffix("returning id").
		ToSql()

	if err != nil {
		return err
	}

	row := r.pool.QueryRow(ctx, sql, v...)

	id := 0
	if err := row.Scan(&id); err != nil {
		return err
	}

	u.Id = id

	return nil
}
