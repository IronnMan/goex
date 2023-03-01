package seeders

import "goex/pkg/seed"

func Initialize() {

	seed.SetRunOrder([]string{
		"SeedUsersTable",
	})
}
