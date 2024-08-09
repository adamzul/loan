package loan

import (
	"context"

	"loan.com/helper/customerr"
)

func (u *loan) GetOutStanding(ctx context.Context, loanID int32) (float64, error) {
	loan, err := u.loanRepo.Get(ctx, loanID)
	if err != nil {
		return 0, customerr.StackTrace(err)
	}

	payments, err := u.paymentRepo.List(ctx, loan.ID)
	if err != nil {
		return 0, customerr.StackTrace(err)
	}

	totalPayment := float64(0)
	for _, payment := range payments {
		totalPayment += payment.Amount
	}
	interestAmount := loan.Amount * loan.Interest / 100
	outStanding := loan.Amount + interestAmount - totalPayment
	return outStanding, nil
}
