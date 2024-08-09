package handlers

import (
	"context"
)

type loan interface {
	GetOutStanding(ctx context.Context, loanID int32) (float64, error)
	IsUserDelinquent(ctx context.Context, loanID int32) (bool, error)
}

type payment interface {
	Pay(ctx context.Context, loanID int32, amount float64) error
}
