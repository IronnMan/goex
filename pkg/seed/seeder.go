package seed

import "gorm.io/gorm"

type SeederFunc func(db *gorm.DB)

type Seeder struct {
	Func SeederFunc
	Name string
}

var seeders []Seeder

var orderedSeederNames []string

func Add(name string, fn SeederFunc) {
	seeders = append(seeders, Seeder{
		Name: name,
		Func: fn,
	})
}

func SetRunOrder(names []string) {
	orderedSeederNames = names
}
