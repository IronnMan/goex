package auth

import (
	"fmt"
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

	// parse JSON request
	if err := c.ShouldBindJSON(&request); err != nil {
		// parsing failed, return 422 status code and error message
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		fmt.Println(err.Error())
		return
	}

	errs := requests.ValidateSignupEmailExist(&request, c)

	if len(errs) > 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}
