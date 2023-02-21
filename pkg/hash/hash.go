package hash

import (
	"goex/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

// BcryptHash encrypt password with bcrypt
func BcryptHash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	logger.LogIf(err)

	return string(bytes)
}

// BcryptCheck compare plaintext password with database hash
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// BcryptIsHashed determine whether a string is hash data
func BcryptIsHashed(str string) bool {
	return len(str) == 60
}
