package models

import "time"

type Payment struct {
	ID        int32     `db:"id" json:"id"`
	ClientID  string    `db:"client_id" json:"client_id"`
	LoanID    string    `db:"loan_id" json:"loan_id"`
	Amount    float64   `db:"amount" json:"amount"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
