package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func (a *application) testEmail(c echo.Context) error {
	bills, err := a.billRepo.GetUpcoming()
	if err != nil {
		return err
	}

	fmt.Println(bills)
	return nil
}
