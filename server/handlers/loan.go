package handlers

import (
	"net/http"

	"loan.com/helper/customerr"

	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type LoanHandler struct {
	v       *validator.Validate
	loanSvc loan
}

func NewLoanHandler(v *validator.Validate, loan loan) *LoanHandler {
	return &LoanHandler{
		v:       v,
		loanSvc: loan,
	}
}

type GetOutStandingReq struct {
	LoanID int32 `json:"loan_id" validate:"required"`
}

type GetOutStadingRes struct {
	Amount float64 `json:"amount"`
}

func (h *LoanHandler) GetOutStanding(c echo.Context) error {
	ctx := c.Request().Context()
	req := GetOutStandingReq{}

	if err := c.Bind(&req); err != nil {
		return customerr.StackTrace(err)
	}

	if err := h.v.StructCtx(ctx, req); err != nil {
		return customerr.StackTrace(ErrorResponse(c, http.StatusBadRequest, "Required fields are empty", nil))
	}

	outStanding, err := h.loanSvc.GetOutStanding(ctx, req.LoanID)
	if err != nil {
		return (ErrorResponse(c, http.StatusInternalServerError, err.Error(), err))
	}

	return Response(c, http.StatusCreated, GetOutStadingRes{
		Amount: outStanding,
	})
}

type IsUserDelinquentReq struct {
	LoanID int32 `json:"loan_id" validate:"required"`
}

type IsUserDelinquentRes struct {
	IsUserDelinquent bool `json:"is_user_delinquent"`
}

func (h *LoanHandler) IsUserDelinquent(c echo.Context) error {
	ctx := c.Request().Context()
	req := IsUserDelinquentReq{}

	if err := c.Bind(&req); err != nil {
		return customerr.StackTrace(err)
	}

	if err := h.v.StructCtx(ctx, req); err != nil {
		return customerr.StackTrace(ErrorResponse(c, http.StatusBadRequest, "Required fields are empty", nil))
	}

	isUserDelinquent, err := h.loanSvc.IsUserDelinquent(ctx, req.LoanID)
	if err != nil {
		return (ErrorResponse(c, http.StatusInternalServerError, err.Error(), err))
	}

	return Response(c, http.StatusCreated, IsUserDelinquentRes{
		IsUserDelinquent: isUserDelinquent,
	})
}
