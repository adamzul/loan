package models

import "time"

type Loan struct {
	ID              int32     `db:"id" json:"id"`
	ClientID        int32     `db:"client_id" json:"client_id"`
	Amount          float64   `db:"amount" json:"amount"`
	Interest        float64   `db:"interest" json:"interest"`
	NumberOfPayment int32     `db:"number_of_payment" json:"number_of_payment"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
}
