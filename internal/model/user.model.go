package model

import "time"

type User struct {
	Id            int       `db:"id"`
	Email         string    `db:"email"`
	Fullname      string    `db:"full_name"`
	Password_hash string    `db:"password_hash"`
	Created_at    time.Time `db:"created_at"`
}
