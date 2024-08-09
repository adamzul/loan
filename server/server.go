package server

import (
	"loan.com/config"
	"loan.com/server/handlers"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo           *echo.Echo
	DB             *gorm.DB
	Config         *config.Config
	LoanHandler    handlers.LoanHandler
	PaymentHandler handlers.PaymentHandler
}

func NewServer(cfg *config.Config, loanHandler handlers.LoanHandler, paymentHandler handlers.PaymentHandler) *Server {
	return &Server{
		Echo:           echo.New(),
		Config:         cfg,
		LoanHandler:    loanHandler,
		PaymentHandler: paymentHandler,
	}
}

func (server *Server) Start(addr string) error {
	return server.Echo.Start(":" + addr)
}
