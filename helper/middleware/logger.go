package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"loan.com/helper/customerr"
)

func LogError() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err != nil {
				err, ok := err.(customerr.Error)
				if ok {
					fmt.Println(err.StackTrace())
				}
			}
			return err
		}
	}
}
