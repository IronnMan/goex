package validators

import (
	"goex/pkg/captcha"
)

// ValidateCaptcha custom rules, verification "Image verification code"
func ValidateCaptcha(captchaID, captchaAnswer string, errs map[string][]string) map[string][]string {
	if ok := captcha.NewCaptcha().VerifyCaptcha(captchaID, captchaAnswer); !ok {
		errs["captcha_answer"] = append(errs["captcha_answer"], "Image verification code")
	}
	return errs
}
