package auth

import (
	"github.com/gin-gonic/gin"
	v1 "goex/app/http/controllers/api/v1"
	"goex/app/models/user"
	"goex/app/requests"
	"net/http"
)

// SignupController register controller
type SignupController struct {
	v1.BaseAPIController
}

func (sc *SignupController) IsEmailExist(c *gin.Context) {
	request := requests.SignupEmailExistRequest{}
	if ok := requests.Validate(c, &request, requests.SignupEmailExist); !ok {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}
