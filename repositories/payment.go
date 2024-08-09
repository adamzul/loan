package repositories

import (
	"context"
	"time"

	"loan.com/models"
	"loan.com/repositories/executor"

	"github.com/huandu/go-sqlbuilder"
)

type payment struct {
	executor *executor.Executor
}

func NewPayment(exec *executor.Executor) *payment {
	return &payment{
		executor: exec,
	}
}

type CreateOpt struct {
	Amount   float64
	LoanID   int32
	ClientID int32
}

func (r *payment) Create(ctx context.Context, opt CreateOpt) error {
	sb := sqlbuilder.NewInsertBuilder()
	sb.InsertInto("payment").Cols(
		"client_id", "loan_id", "amount", "created_at").
		Values(opt.ClientID, opt.LoanID, opt.Amount, time.Now())

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	_, err := r.executor.ExecContext(ctx, query, args...)
	return err
}

func (r *payment) List(ctx context.Context, loanID int32) ([]models.Payment, error) {
	whereCondition := []string{}
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("*").From("payment")

	whereCondition = append(whereCondition,
		sb.Equal("loan_id", loanID),
	)

	sb.Where(
		whereCondition...,
	)

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	data := []models.Payment{}
	err := r.executor.SelectContext(ctx, &data, query, args...)
	return data, err
}

func (r *payment) Count(ctx context.Context, loanID int32) (int32, error) {
	whereCondition := []string{}
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("count(*)").From("payment")

	whereCondition = append(whereCondition,
		sb.Equal("loan_id", loanID),
	)

	sb.Where(
		whereCondition...,
	)

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)
	count := int32(0)
	err := r.executor.GetContext(ctx, &count, query, args...)
	return count, err
}
