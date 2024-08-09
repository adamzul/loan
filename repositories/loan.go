package repositories

import (
	"context"

	"loan.com/models"
	"loan.com/repositories/executor"

	"github.com/huandu/go-sqlbuilder"
)

type loan struct {
	executor *executor.Executor
}

func NewLoan(exec *executor.Executor) *loan {
	return &loan{
		executor: exec,
	}
}

func (r *loan) Get(ctx context.Context, loanID int32) (models.Loan, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("*").From("loan").Where(
		sb.Equal("id", loanID),
	)
	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	data := models.Loan{}
	err := r.executor.GetContext(ctx, &data, query, args...)
	return data, err
}
