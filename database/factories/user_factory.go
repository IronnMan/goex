package factories

import (
	"github.com/go-faker/faker/v4"
	"goex/app/models/user"
)

func MakeUsers(times int) []user.User {

	var objs []user.User

	faker.SetGenerateUniqueValues(true)

	for i := 0; i < times; i++ {
		model := user.User{
			Name:     faker.Username(),
			Email:    faker.Email(),
			Password: "$2a$14$JoyAAZpWdxMF61IloG4IGumvvgsYowkNCkeMOS20dP73y8ZqjfZJi",
		}
		objs = append(objs, model)
	}

	return objs
}
