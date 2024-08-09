package loan

import (
	"context"

	"loan.com/models"
)

type loanRepo interface {
	Get(ctx context.Context, loanID int32) (models.Loan, error)
}

type paymentRepo interface {
	List(ctx context.Context, loanID int32) ([]models.Payment, error)
	Count(ctx context.Context, loanID int32) (int32, error)
}

type accountRepo interface {
	GetByAccNo(ctx context.Context, clientID string) (models.Account, error)
}
