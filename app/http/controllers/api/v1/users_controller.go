package v1

import (
	"github.com/gin-gonic/gin"
	"goex/pkg/auth"
	"goex/pkg/response"
)

type UsersController struct {
	BaseAPIController
}

func (ctrl *UsersController) CurrentUser(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	response.Data(c, userModel)
}
