package mail

import (
	"fmt"
	emailPKG "github.com/jordan-wright/email"
	"goex/pkg/logger"
	"net/smtp"
)

// SMTP implement email.Driver interface
type SMTP struct{}

// Send implement the Send method of the email.Driver interface
func (s *SMTP) Send(email Email, config map[string]string) bool {

	e := emailPKG.NewEmail()

	e.From = fmt.Sprintf("%v <%v>", email.From.Name, email.From.Address)
	e.To = email.To
	e.Bcc = email.Bcc
	e.Cc = email.Cc
	e.Subject = email.Subject
	e.Text = email.Text
	e.HTML = email.HTML

	logger.DebugJSON("Send Email", "Shipping details", e)

	err := e.Send(
		fmt.Sprintf("%v:%v", config["host"], config["port"]),

		smtp.PlainAuth(
			"",
			config["username"],
			config["password"],
			config["host"],
		),
	)
	if err != nil {
		logger.ErrorString("Send Email", "Sending error", err.Error())
		return false
	}

	logger.DebugString("Send Email", "Sending successfully", "")
	return true
}
