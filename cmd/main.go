package main

import (
	application "loan.com"

	"loan.com/config"
)

func main() {
	cfg := config.MustLoad()

	application.Start(cfg)
}
