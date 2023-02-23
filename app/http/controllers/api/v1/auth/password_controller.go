package auth

import (
	"github.com/gin-gonic/gin"
	v1 "goex/app/http/controllers/api/v1"
	"goex/app/models/user"
	"goex/app/requests"
	"goex/pkg/response"
)

type PasswordController struct {
	v1.BaseAPIController
}

func (pc *PasswordController) ResetByEmail(c *gin.Context) {
	request := requests.ResetByEmailRequest{}
	if ok := requests.Validate(c, &request, requests.ResetByEmail); !ok {
		return
	}

	userModel := user.GetByEmail(request.Email)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()
		response.Success(c)
	}
}
