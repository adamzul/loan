package application

import (
	"log"

	"github.com/go-playground/validator/v10"
	"loan.com/config"
	"loan.com/connection"
	"loan.com/repositories"
	"loan.com/repositories/executor"
	"loan.com/server"
	"loan.com/server/handlers"
	"loan.com/server/routes"
	"loan.com/services/loan"
	"loan.com/services/payment"
)

func Start(cfg *config.Config) {
	replicationDB := connection.ReplicationDB{
		Primary: connection.NewPostgresSQL(&cfg.Postgres.Primary, "pgx"),
		Standby: connection.NewPostgresSQL(&cfg.Postgres.Standby, "pgx"),
	}

	executorRepo := executor.New(&replicationDB)
	transactionRepo := executor.NewTransaction(&replicationDB)

	accountRepo := repositories.NewAccount(&executorRepo)
	paymentRepo := repositories.NewPayment(&executorRepo)
	loanRepo := repositories.NewLoan(&executorRepo)

	loanSvc := loan.New(transactionRepo, accountRepo, loanRepo, paymentRepo)
	paymentSvc := payment.New(transactionRepo, loanRepo, paymentRepo)

	balanceHandler := handlers.NewLoanHandler(validator.New(), loanSvc)
	paymentHandler := handlers.NewPaymentHandler(validator.New(), paymentSvc)

	app := server.NewServer(cfg, *balanceHandler, *paymentHandler)

	routes.ConfigureRoutes(app)

	err := app.Start(cfg.HTTP.Port)
	if err != nil {
		log.Fatal("Port already used")
	}
}
