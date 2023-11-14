package main

import (
	"bytes"
	"context"
	"fmt"

	"github.com/nathanaelcunningham/billReminder/components"
)

func (a *application) RunEmailCron() {
	bills, err := a.billRepo.GetUpcoming()
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	content := components.BillReminderEmail(bills)
	buf := bytes.NewBuffer(nil)
	content.Render(context.Background(), buf)
	err = a.mailClient.SendMail(buf.String())
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}
