package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"goex/pkg/response"
)

// ValidatorFunc validation function type
type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

// Validate example call in controller:
//
//	if ok := requests.Validate(c, &requests.UserSaveRequest{}, requests.UserSave); !ok {
//	    return
//	}
func Validate(c *gin.Context, obj interface{}, handler ValidatorFunc) bool {
	// parse requests, support JSON data, form requests and URL query
	if err := c.ShouldBind(obj); err != nil {
		response.BadRequest(c, err, "Request parsing error, please confirm whether the request format is correct. Please use the multipart header when uploading files, and use JSON format for parameters.")
		return false
	}

	errs := handler(obj, c)

	if len(errs) > 0 {
		response.ValidationError(c, errs)
		return false
	}

	return true
}

func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	// Configuration options
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid", // struct tag identifier in the model
		Messages:      messages,
	}

	return govalidator.New(opts).ValidateStruct()
}
