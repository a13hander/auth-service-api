package model

import (
	"database/sql"
	"time"
)

type Access struct {
	Id       int          `db:"id"`
	Endpoint string       `db:"endpoint"`
	Role     int          `db:"role"`
	Created  time.Time    `db:"created_at"`
	Updated  sql.NullTime `db:"updated_at"`
}
