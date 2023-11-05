package main

import (
	"bytes"

	"github.com/labstack/echo/v4"
	"github.com/nathanaelcunningham/billReminder/components"
)

func (a *application) home(c echo.Context) error {
	bills, err := a.billRepo.GetAll()
	if err != nil {
		return err
	}

	comp := components.Layout(bills)
	buf := bytes.NewBuffer(nil)
	comp.Render(c.Request().Context(), buf)
	return c.HTML(200, buf.String())
}
