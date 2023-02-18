package mail

type Driver interface {
	// Send check captcha
	Send(email Email, config map[string]string) bool
}
