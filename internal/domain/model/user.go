package model

import "time"

type User struct {
	Id        int
	Email     string
	Username  string
	Password  string
	Role      int
	CreatedAt time.Time
	UpdatedAt time.Time
}
