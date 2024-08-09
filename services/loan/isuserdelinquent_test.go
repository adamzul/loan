package loan_test

import (
	"context"
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"loan.com/models"
	"loan.com/services/loan"
	"loan.com/services/payment/mock"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	BeforeSuite(func() {})
	RunSpecs(t, "")
}

var _ = Describe("IsUserDelinquent", func() {
	type usecase interface {
		IsUserDelinquent(ctx context.Context, loanID int32) (bool, error)
	}
	var (
		paymentRepo *mock.MockpaymentRepo
		loanRepo    *mock.MockloanRepo
		clockMock   *clock.Mock
		mockCtrl    *gomock.Controller

		uc usecase
	)

	BeforeEach(func() {
		t := GinkgoT()

		mockCtrl = gomock.NewController(t)
		paymentRepo = mock.NewMockpaymentRepo(mockCtrl)
		loanRepo = mock.NewMockloanRepo(mockCtrl)
		clockMock = clock.NewMock()
		now, _ := time.Parse(time.DateOnly, "2024-08-08")
		clockMock.Set(now)
		uc = loan.New(loanRepo, paymentRepo, clockMock)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	ctx := context.TODO()

	Context("OK", func() {
		It("success and user not delinquent", func() {
			loanID := int32(1)
			loanCreatedAt, _ := time.Parse(time.DateOnly, "2024-07-08")
			loanRepo.EXPECT().Get(ctx, loanID).Return(models.Loan{
				ID:              loanID,
				ClientID:        1,
				Amount:          5_000_000,
				Interest:        10,
				NumberOfPayment: 50,
				CreatedAt:       loanCreatedAt,
			}, nil)

			paymentRepo.EXPECT().Count(ctx, loanID).Return(int32(4), nil)

			isDelinquent, err := uc.IsUserDelinquent(ctx, loanID)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(isDelinquent).To(BeEquivalentTo(false))
		})

	})

})
