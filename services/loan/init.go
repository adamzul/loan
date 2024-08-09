package loan

import (
	"loan.com/repositories/executor"
)

type loan struct {
	trx         executor.Transaction
	accountRepo accountRepo
	paymentRepo paymentRepo
	loanRepo    loanRepo
}

func New(trx executor.Transaction, accountRepo accountRepo, loanRepo loanRepo, paymentRepo paymentRepo) *loan {
	return &loan{
		trx:         trx,
		accountRepo: accountRepo,
		paymentRepo: paymentRepo,
		loanRepo:    loanRepo,
	}
}
