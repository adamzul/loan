package payment_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"loan.com/models"
	"loan.com/repositories"
	"loan.com/services/payment"
	"loan.com/services/payment/mock"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	BeforeSuite(func() {})
	RunSpecs(t, "")
}

var _ = Describe("Pay", func() {
	type usecase interface {
		Pay(ctx context.Context, loanID int32, amount float64) error
	}
	var (
		paymentRepo *mock.MockpaymentRepo
		loanRepo    *mock.MockloanRepo
		trx         *mock.Mocktransaction

		mockCtrl *gomock.Controller

		uc usecase
	)

	BeforeEach(func() {
		t := GinkgoT()

		mockCtrl = gomock.NewController(t)
		paymentRepo = mock.NewMockpaymentRepo(mockCtrl)
		loanRepo = mock.NewMockloanRepo(mockCtrl)
		trx = mock.NewMocktransaction(mockCtrl)

		uc = payment.New(trx, loanRepo, paymentRepo)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	ctx := context.TODO()

	Context("OK", func() {
		It("success create payment", func() {
			loanID := int32(1)
			amount := float64(110_000)
			loanCreatedAt, _ := time.Parse(time.DateOnly, "2024-07-08")
			trx.EXPECT().Execute(ctx, gomock.Any()).DoAndReturn(MockTransactionCall)
			loanRepo.EXPECT().Get(ctx, loanID).Return(models.Loan{
				ID:              loanID,
				ClientID:        1,
				Amount:          5_000_000,
				Interest:        10,
				NumberOfPayment: 50,
				CreatedAt:       loanCreatedAt,
			}, nil)

			paymentRepo.EXPECT().List(ctx, loanID).Return(nil, nil)
			paymentRepo.EXPECT().Create(ctx, repositories.CreateOpt{
				Amount:   amount,
				LoanID:   loanID,
				ClientID: 1,
			})
			err := uc.Pay(ctx, loanID, amount)
			Expect(err).ShouldNot(HaveOccurred())
		})

	})

})

func MockTransactionCall(ctx context.Context, fn func(context.Context) error) error {
	var err error
	func() {
		defer func() {
			if p := recover(); p != nil {
				switch e := p.(type) {
				case error:
					err = e
				default:
					panic(e)
				}
			}
		}()
		err = fn(ctx)
	}()

	return err
}
