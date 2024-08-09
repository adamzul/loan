package loan

import "github.com/benbjohnson/clock"

type loan struct {
	paymentRepo paymentRepo
	loanRepo    loanRepo
	clock       clock.Clock
}

func New(loanRepo loanRepo, paymentRepo paymentRepo, clock clock.Clock) *loan {
	return &loan{
		paymentRepo: paymentRepo,
		loanRepo:    loanRepo,
		clock:       clock,
	}
}
