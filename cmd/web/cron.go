package main

import (
	"bytes"
	"context"

	"log/slog"

	"github.com/nathanaelcunningham/billReminder/components"
)

func (a *application) RunEmailCron() {
	bills, err := a.billRepo.GetAll()
	if err != nil {
		slog.Error(err.Error())
	}

	content := components.BillReminderEmail(bills)
	buf := bytes.NewBuffer(nil)
	content.Render(context.Background(), buf)
	err = a.mailClient.SendMail(buf.String())
	if err != nil {
		slog.Error(err.Error())
	}
}
