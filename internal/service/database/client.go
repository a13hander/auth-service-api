package database

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var _ Client = (*client)(nil)

type Query struct {
	Name     string
	QueryRaw string
}

type QueryExecer interface {
	ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row
}

type NamedExecer interface {
	Get(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
	Select(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
}

type Client interface {
	QueryExecer
	NamedExecer
	Close() error
}

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type client struct {
	pg *pgxpool.Pool
}

func NewClient(ctx context.Context, conf DbConfig) Client {
	return &client{
		pg: createPgxPool(ctx, conf),
	}
}

func (c *client) ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error) {
	return c.pg.Exec(ctx, q.QueryRaw, args...)
}

func (c *client) QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error) {
	return c.pg.Query(ctx, q.QueryRaw, args...)
}

func (c *client) QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row {
	return c.pg.QueryRow(ctx, q.QueryRaw, args...)
}

func (c *client) Get(ctx context.Context, dest interface{}, q Query, args ...interface{}) error {
	rows, err := c.QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanOne(dest, rows)
}

func (c *client) Select(ctx context.Context, dest interface{}, q Query, args ...interface{}) error {
	rows, err := c.QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanAll(dest, rows)
}

func (c *client) Close() error {
	c.pg.Close()
	return nil
}
