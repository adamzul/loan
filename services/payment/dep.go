package payment

import (
	"context"

	"loan.com/models"
	"loan.com/repositories"
)

type paymentRepo interface {
	Create(ctx context.Context, opt repositories.CreateOpt) error
	List(ctx context.Context, loanID int32) ([]models.Payment, error)
	Count(ctx context.Context, LoanID int32) (int32, error)
}

type loanRepo interface {
	Get(ctx context.Context, loanID int32) (models.Loan, error)
}
