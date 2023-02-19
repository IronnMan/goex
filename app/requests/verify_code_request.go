package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"goex/app/requests/validators"
)

type VerifyCodeEmailRequest struct {
	CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`

	Email string `json:"email,omitempty" valid:"email"`
}

func VerifyCodeEmail(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"email":          []string{"required", "min:4", "max:30", "email"},
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"email": []string{
			"required:Email is required",
			"min:Email length must be greater than 4",
			"max:Email length must be less than 30",
			"email:Email format is incorrect, please provide a valid email address",
		},
		"captcha_id": []string{
			"required:The ID of the image captcha is required",
		},
		"captcha_answer": []string{
			"required:Image verification code answer is required",
			"digits:Image captcha must be 6 digits long",
		},
	}

	errs := validate(data, rules, messages)

	_data := data.(*VerifyCodeEmailRequest)
	errs = validators.ValidateCaptcha(_data.CaptchaID, _data.CaptchaAnswer, errs)

	return errs
}
