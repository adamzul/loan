package routes

import (
	custommiddleware "loan.com/helper/middleware"
	s "loan.com/server"

	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func ConfigureRoutes(server *s.Server) {
	server.Echo.Use(middleware.Logger())

	server.Echo.GET("/swagger/*", echoSwagger.WrapHandler)

	r := server.Echo.Group("")

	r.Use(custommiddleware.LogError())
	r.POST("/loan/out-standing", server.LoanHandler.GetOutStanding)
	r.POST("/loan/is-user-delinquent", server.LoanHandler.IsUserDelinquent)
	r.POST("/payment/pay", server.PaymentHandler.Pay)
}
