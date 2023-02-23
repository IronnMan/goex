package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"goex/app/models/user"
	"goex/pkg/logger"
)

// Attempt log in
func Attempt(email string, password string) (user.User, error) {
	userModel := user.GetByMulti(email)
	if userModel.ID == 0 {
		return user.User{}, errors.New("account does not exist")
	}

	if !userModel.ComparePassword(password) {
		return user.User{}, errors.New("wrong password")
	}

	return userModel, nil
}

// CurrentUser get currently logged in user from gin.context
func CurrentUser(c *gin.Context) user.User {
	userModel, ok := c.MustGet("current_user").(user.User)
	if !ok {
		logger.LogIf(errors.New("could not ger user"))
		return user.User{}
	}

	return userModel
}

// CurrentUID get currently logged-in user ID from gin.context
func CurrentUID(c *gin.Context) string {
	return c.GetString("current_user_id")
}
