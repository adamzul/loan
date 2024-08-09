package payment

import (
	"loan.com/repositories/executor"
)

type loan struct {
	trx         executor.Transaction
	paymentRepo paymentRepo
	loanRepo    loanRepo
}

func New(trx executor.Transaction, loanRepo loanRepo, paymentRepo paymentRepo) *loan {
	return &loan{
		trx:         trx,
		loanRepo:    loanRepo,
		paymentRepo: paymentRepo,
	}
}
