package models

import "time"

type Account struct {
	ID        int32     `db:"id"`
	Username  string    `db:"username"`
	CreatedAt time.Time `db:"created_at"`
}
