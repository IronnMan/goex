package user

import (
	"goex/app/models"
	"goex/pkg/database"
)

// User model
type User struct {
	models.BaseModel

	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}

func (u *User) Create() {
	database.DB.Create(&u)
}
