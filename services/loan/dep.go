package loan

import (
	"context"

	"loan.com/models"
)

//go:generate mockgen -source=dep.go -destination=mock/mock.go -package=mock
type loanRepo interface {
	Get(ctx context.Context, loanID int32) (models.Loan, error)
}

type paymentRepo interface {
	List(ctx context.Context, loanID int32) ([]models.Payment, error)
	Count(ctx context.Context, loanID int32) (int32, error)
}
