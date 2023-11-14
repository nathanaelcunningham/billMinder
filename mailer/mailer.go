package mailer

import (
	"fmt"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Mailer struct {
	client *sendgrid.Client
}

func NewMailer(apiKey string) *Mailer {
	client := sendgrid.NewSendClient(apiKey)
	return &Mailer{
		client: client,
	}
}

func (m Mailer) SendMail(content string) error {
	from := mail.NewEmail("Bill Minder", "billminder@thinkbrain.dev")
	to := mail.NewEmail("Nathanael Cunningham", "nathanaelcunningham@gmail.com")

	message := mail.NewSingleEmail(from, "Weekly Bill List", to, "", content)
	response, err := m.client.Send(message)
	if err != nil {
		return err
	}
	if response.StatusCode >= 400 {
		return fmt.Errorf("error sending email: %v", response.Body)
	}
	fmt.Println("Email sent successfully")
	return nil
}
