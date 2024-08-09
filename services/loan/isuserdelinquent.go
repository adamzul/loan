package loan

import (
	"context"
	"math"
	"time"

	"loan.com/helper/customerr"
)

func (u *loan) IsUserDelinquent(ctx context.Context, loanID int32) (bool, error) {
	loan, err := u.loanRepo.Get(ctx, loanID)
	if err != nil {
		return false, customerr.StackTrace(err)
	}

	count, err := u.paymentRepo.Count(ctx, loanID)
	if err != nil {
		return false, customerr.StackTrace(err)
	}

	now := time.Now()
	durationDay := now.Sub(loan.CreatedAt).Hours() / 24
	durationWeek := int(math.Ceil(durationDay / 7))
	if durationWeek-int(count) >= 2 {
		return true, nil
	}
	return false, nil
}
