package repositories

import (
	"context"

	"loan.com/models"
	"loan.com/repositories/executor"

	"github.com/huandu/go-sqlbuilder"
)

type account struct {
	executor *executor.Executor
}

func NewAccount(exec *executor.Executor) *account {
	return &account{
		executor: exec,
	}
}

func (r *account) GetByAccNo(ctx context.Context, accountNumber string) (models.Account, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("*").From("account").Where(
		sb.Equal("account_number", accountNumber),
	)
	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	data := models.Account{}
	err := r.executor.GetContext(ctx, &data, query, args...)
	return data, err
}
