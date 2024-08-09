package handlers

import (
	"net/http"

	"loan.com/helper/customerr"

	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type PaymentHandler struct {
	v          *validator.Validate
	paymentSvc payment
}

func NewPaymentHandler(v *validator.Validate, payment payment) *PaymentHandler {
	return &PaymentHandler{
		v:          v,
		paymentSvc: payment,
	}
}

type PayReq struct {
	LoanID int32   `json:"loan_id" validate:"required"`
	Amount float64 `json:"amount" validate:"required"`
}

func (h *PaymentHandler) Pay(c echo.Context) error {
	ctx := c.Request().Context()
	req := PayReq{}

	if err := c.Bind(&req); err != nil {
		return customerr.StackTrace(err)
	}

	if err := h.v.StructCtx(ctx, req); err != nil {
		return customerr.StackTrace(ErrorResponse(c, http.StatusBadRequest, "Required fields are empty", nil))
	}

	err := h.paymentSvc.Pay(ctx, req.LoanID, req.Amount)
	if err != nil {
		return (ErrorResponse(c, http.StatusInternalServerError, err.Error(), err))
	}

	return MessageResponse(c, http.StatusCreated, "payment success")
}
