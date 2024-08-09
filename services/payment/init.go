package payment

type payment struct {
	trx         transaction
	paymentRepo paymentRepo
	loanRepo    loanRepo
}

func New(trx transaction, loanRepo loanRepo, paymentRepo paymentRepo) *payment {
	return &payment{
		trx:         trx,
		loanRepo:    loanRepo,
		paymentRepo: paymentRepo,
	}
}
