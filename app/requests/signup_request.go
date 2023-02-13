package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

func SignupEmailExist(data interface{}, c *gin.Context) map[string][]string {

	// custom validation rules
	rules := govalidator.MapData{
		"email": []string{"required", "min:4", "max:30", "email"},
	}

	// prompt when custom validation error occurs
	messages := govalidator.MapData{
		"email": []string{
			"required:Email is required",
			"min:Email length must be greater than 4",
			"max:Email length must be less than 30",
			"email:Email format is incorrect, please provide a valid email address",
		},
	}

	return validate(data, rules, messages)
}
