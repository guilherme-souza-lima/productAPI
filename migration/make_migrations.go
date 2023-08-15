package migration

import (
	"gorm.io/gorm"
	"productAPI/migration/model"
)

type DatabaseMakeMigrations struct {
	postgres *gorm.DB
}

func NewDatabaseMakeMigrations(postgres *gorm.DB) DatabaseMakeMigrations {
	return DatabaseMakeMigrations{postgres}
}

func (ref DatabaseMakeMigrations) MakeMigrations() {
	ref.makeMigrationsDatabase()
}

func (ref DatabaseMakeMigrations) makeMigrationsDatabase() {
	err := ref.postgres.Migrator().AutoMigrate(&model.Product{})
	errFunc(err)
	err = ref.postgres.Migrator().AutoMigrate(&model.Transaction{})
	errFunc(err)
}

func errFunc(err error) {
	if err != nil {
		panic("Error Migrating: " + err.Error())
	}
}
