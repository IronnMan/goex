package auth

import (
	"errors"
	"goex/app/models/user"
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
