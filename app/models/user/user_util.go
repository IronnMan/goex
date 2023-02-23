package user

import "goex/pkg/database"

// IsEmailExist judging that Email has been registered
func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// GetByMulti get user by email/name
func GetByMulti(loginID string) (userModel User) {
	database.DB.Where("email = ?", loginID).Or("name = ?", loginID).First(&userModel)
	return
}

func Get(idstr string) (userModel User) {
	database.DB.Where("id", idstr).First(&userModel)
	return
}

func GetByEmail(email string) (userModel User) {
	database.DB.Where("email = ?", email).First(&userModel)
	return
}
