package main

import (
	"bytes"

	"github.com/labstack/echo/v4"
	"github.com/nathanaelcunningham/billReminder/components"
)

func (a *application) testEmail(c echo.Context) error {
	bills, err := a.billRepo.GetAll()
	if err != nil {
		return err
	}

	content := components.BillReminderEmail(bills)
	buf := bytes.NewBuffer(nil)
	content.Render(c.Request().Context(), buf)
	err = a.mailClient.SendMail(buf.String())
	if err != nil {
		return err
	}
	return nil
}
