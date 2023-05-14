package database

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"github.com/a13hander/auth-service-api/internal/domain/model"
)

const tableName = "users"

type UserRepo struct {
	dbClient Client
}

func NewUserRepo(dbClient Client) *UserRepo {
	return &UserRepo{
		dbClient: dbClient,
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

	q := Query{
		Name:     "UserRepo.Create",
		QueryRaw: sql,
	}
	row := r.dbClient.QueryRowContext(ctx, q, v...)

	id := 0
	if err := row.Scan(&id); err != nil {
		return err
	}

	u.Id = id

	return nil
}

func (r *UserRepo) GetAll(ctx context.Context) ([]*model.User, error) {
	sql, args, err := sq.Select("id", "email", "username", "role", "created_at").
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	q := Query{
		Name:     "UserRepo.GetAll",
		QueryRaw: sql,
	}

	row, err := r.dbClient.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	users := make([]*model.User, 0, 100)
	for row.Next() {
		u := &model.User{}
		err := row.Scan(&u.Id, &u.Email, &u.Username, &u.Role, &u.CreatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}
