package verifycode

import (
	"fmt"
	"goex/pkg/app"
	"goex/pkg/config"
	"goex/pkg/helpers"
	"goex/pkg/logger"
	"goex/pkg/mail"
	"goex/pkg/redis"
	"strings"
	"sync"
)

type VerifyCode struct {
	Store Store
}

var once sync.Once
var internalVerifyCode *VerifyCode

// NewVerifyCode singleton mode acquisition
func NewVerifyCode() *VerifyCode {
	once.Do(func() {
		internalVerifyCode = &VerifyCode{
			Store: &RedisStore{
				RedisClient: redis.Redis,
				KeyPrefix:   config.GetString("app.name") + ":verifycode:",
			},
		}
	})

	return internalVerifyCode
}

func (vc *VerifyCode) SendEmail(email string) error {

	code := vc.generateVerifyCode(email)

	if !app.IsProduction() && strings.HasSuffix(email, config.GetString("verifycode.debug_email_suffix")) {
		return nil
	}

	content := fmt.Sprintf("<h1>Your Email verification code is %v</h1>", code)

	mail.NewMailer().Send(mail.Email{
		From: mail.From{
			Address: config.GetString("mail.from.address"),
			Name:    config.GetString("mail.from.name"),
		},
		To:      []string{email},
		Subject: "Email verification code",
		HTML:    []byte(content),
	})

	return nil
}

// CheckAnswer check whether the verification code submitted by the user is correct
func (vc *VerifyCode) CheckAnswer(key string, answer string) bool {
	logger.DebugJSON("Verification code", "Check verification code", map[string]string{key: answer})

	if !app.IsProduction() && strings.HasSuffix(key, config.GetString("verifycode.debug_email_suffix")) {
		return true
	}

	return vc.Store.Verify(key, answer, false)
}

// generateVerifyCode generate a verification code and place it in Redis
func (vc *VerifyCode) generateVerifyCode(key string) string {

	// generate random code
	code := helpers.RandomNumber(config.GetInt("verifycode.code_length"))

	if app.IsLocal() {
		code = config.GetString("verifycode.debug_code")
	}

	logger.DebugJSON("Verification code", "Generate verification code", map[string]string{key: code})

	// store the verification code and KEY (email or other) in Redis and set the expiration time
	vc.Store.Set(key, code)
	return code
}
