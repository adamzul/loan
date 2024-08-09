package payment

import (
	"context"
	"errors"
	"math"
	"time"

	"loan.com/helper/customerr"
	"loan.com/repositories"
)

func (u *payment) Pay(ctx context.Context, loanID int32, amount float64) error {
	loan, err := u.loanRepo.Get(ctx, loanID)
	if err != nil {
		return customerr.StackTrace(err)
	}
	err = u.trx.Execute(ctx, func(ctx context.Context) error {
		payments, errTrx := u.paymentRepo.List(ctx, loanID)
		if errTrx != nil {
			return customerr.StackTrace(errTrx)
		}

		now := time.Now()
		durationDay := now.Sub(loan.CreatedAt).Hours() / 24
		durationWeek := int(math.Ceil(durationDay / 7))

		missPayment := durationWeek - len(payments)
		if missPayment == 0 {
			return customerr.StackTrace(errors.New("payment to early"))
		}
		lastPaymentDate, _ := time.Parse(time.DateOnly, "1970-01-01")
		if len(payments) != 0 {
			lastPayment := payments[len(payments)-1]
			lastPaymentDate, errTrx = time.Parse(time.DateOnly, lastPayment.CreatedAt.Format(time.DateOnly))
			if errTrx != nil {
				return customerr.StackTrace(errTrx)
			}
		}
		weekDuration := time.Hour * 24 * 7
		if now.Sub(lastPaymentDate) < weekDuration && missPayment < 2 {
			return customerr.StackTrace(errors.New("payment to early"))
		}

		interestAmount := loan.Amount * loan.Interest / 100
		requiredAmount := (loan.Amount + interestAmount) / float64(loan.NumberOfPayment)

		if amount != requiredAmount {
			return customerr.StackTrace(errors.New("amount not match"))
		}

		errTrx = u.paymentRepo.Create(ctx, repositories.CreateOpt{
			Amount:   amount,
			LoanID:   loanID,
			ClientID: loan.ClientID,
		})
		if errTrx != nil {
			return customerr.StackTrace(errTrx)
		}

		return nil
	})

	if err != nil {
		return customerr.StackTrace(err)
	}
	return nil
}
