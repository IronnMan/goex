package user

import "goex/pkg/database"

// IsEmailExist judging that Email has been registered
func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email = ?", email).Count(&count)
	return count > 0
}
