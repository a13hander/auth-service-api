package model

import "time"

type User struct {
	Id             int       `db:"id"`
	Email          string    `db:"email"`
	Username       string    `db:"username"`
	Password       string    `db:"password"`
	Role           int       `db:"role"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
	Specialisation []byte    `db:"specialisation"`
}
