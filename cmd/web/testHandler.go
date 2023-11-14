package main

import (
	"fmt"
	"log/slog"

	"github.com/labstack/echo/v4"
)

func (a *application) testEmail(c echo.Context) error {
	bills, err := a.billRepo.GetUpcoming()
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	fmt.Println(bills)
	return nil
}
