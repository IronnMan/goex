package captcha

import (
	"github.com/mojocn/base64Captcha"
	"goex/pkg/app"
	"goex/pkg/config"
	"goex/pkg/redis"
	"sync"
)

type Captcha struct {
	Base64Captcha *base64Captcha.Captcha
}

// once make sure the internalCaptcha object is only initialized once
var once sync.Once

// internalCaptcha Captcha object used internally
var internalCaptcha *Captcha

func NewCaptcha() *Captcha {
	once.Do(func() {
		// init Captcha object
		internalCaptcha = &Captcha{}

		// use the global Redis object and configure the prefix to store the Key
		store := RedisStore{
			RedisClient: redis.Redis,
			KeyPrefix:   config.GetString("app.name") + ":captcha",
		}

		// Configure base64Captcha driver information
		driver := base64Captcha.NewDriverDigit(
			config.GetInt("captcha.height"),
			config.GetInt("captcha.width"),
			config.GetInt("captcha.length"),
			config.GetFloat64("captcha.maxskew"),
			config.GetInt("captcha.dotcount"),
		)

		internalCaptcha.Base64Captcha = base64Captcha.NewCaptcha(driver, &store)
	})

	return internalCaptcha
}

// GenerateCaptcha generate image verification code
func (c *Captcha) GenerateCaptcha() (id string, b64s string, err error) {
	return c.Base64Captcha.Generate()
}

// VerifyCaptcha verify that the verification code is correct
func (c *Captcha) VerifyCaptcha(id string, answer string) (match bool) {
	if !app.IsProduction() && id == config.GetString("captcha.testing_key") {
		return true
	}

	return c.Base64Captcha.Verify(id, answer, false)
}
