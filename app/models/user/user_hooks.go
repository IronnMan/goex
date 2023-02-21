package user

import (
	"goex/pkg/hash"
	"gorm.io/gorm"
)

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if !hash.BcryptIsHashed(u.Password) {
		u.Password = hash.BcryptHash(u.Password)
	}
	return
}
