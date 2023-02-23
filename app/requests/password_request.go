package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"goex/app/requests/validators"
)

type ResetByEmailRequest struct {
	Email      string `json:"email,omitempty" valid:"email"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
	Password   string `json:"password,omitempty" valid:"password"`
}

func ResetByEmail(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"email":       []string{"required", "min:4", "max:30", "email"},
		"verify_code": []string{"required", "digits:6"},
		"password":    []string{"required", "min:6"},
	}

	messages := govalidator.MapData{}

	errs := validate(data, rules, messages)

	_data := data.(*ResetByEmailRequest)
	errs = validators.ValidateVerifyCode(_data.Email, _data.VerifyCode, errs)

	return errs
}