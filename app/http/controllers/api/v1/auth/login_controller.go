package auth

import (
	"github.com/gin-gonic/gin"
	v1 "goex/app/http/controllers/api/v1"
	"goex/app/requests"
	"goex/pkg/auth"
	"goex/pkg/jwt"
	"goex/pkg/response"
)

type LoginController struct {
	v1.BaseAPIController
}

func (lc *LoginController) LoginByPassword(c *gin.Context) {
	request := requests.LoginByPasswordRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPassword); !ok {
		return
	}

	user, err := auth.Attempt(request.LoginID, request.Password)
	if err != nil {
		response.Unauthorized(c, "The account does not exist or the password is wrong")
	} else {
		token := jwt.NewJWT().IssueToken(jwt.UserInfo{
			UserID:   user.GetStringID(),
			UserName: user.Name,
		})

		response.JSON(c, gin.H{
			"token": token,
		})
	}
}
