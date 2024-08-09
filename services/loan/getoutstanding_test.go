package loan_test

import (
	"context"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"loan.com/models"
	"loan.com/services/loan"
	"loan.com/services/payment/mock"
)

var _ = Describe("GetOutStanding", func() {
	type usecase interface {
		GetOutStanding(ctx context.Context, loanID int32) (float64, error)
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
		clockMock.Set(time.Now())
		uc = loan.New(loanRepo, paymentRepo, clockMock)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	ctx := context.TODO()

	Context("OK", func() {
		It("success and get outstanding", func() {
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

			paymentRepo.EXPECT().List(ctx, loanID).Return(nil, nil)

			outstanding, err := uc.GetOutStanding(ctx, loanID)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(outstanding).To(BeEquivalentTo(5_500_000))
		})

	})

})
