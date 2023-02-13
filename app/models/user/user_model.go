package user

import "goex/app/models"

// User model
type User struct {
	models.BaseModel

	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}
