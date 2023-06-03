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
		Columns("email", "username", "password", "role", "created_at", "updated_at", "specialisation").
		Values(u.Email, u.Username, u.Password, u.Role, u.CreatedAt, u.UpdatedAt, u.Specialisation).
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

func (r *UserRepo) Get(ctx context.Context, username string) (*model.User, error) {
	sql, args, err := sq.Select("id", "email", "username", "password", "role", "created_at", "specialisation").
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"username": username}).
		Limit(1).
		ToSql()

	if err != nil {
		return nil, err
	}

	q := Query{
		Name:     "UserRepo.Get",
		QueryRaw: sql,
	}

	user := &model.User{}
	err = r.dbClient.Get(ctx, user, q, args...)
	if err != nil {
		return nil, err
	}

	return user, nil
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

	const predefinedSize = 100
	users := make([]*model.User, 0, predefinedSize)
	err = r.dbClient.Select(ctx, &users, q, args...)
	if err != nil {
		return nil, err
	}

	return users, nil
}
