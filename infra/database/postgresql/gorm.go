package postgresql

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"productAPI/infra/database"
)

func InitGorm(config *database.Config) *gorm.DB {

	dsn := "host=" + config.Hostname + " user=" + config.UserName + " password=" + config.Password + " dbname=" +
		config.Database + " port=" + config.Port + " sslmode=disable TimeZone=UTC"

	if config.Environment == "integration-test" {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect postgres database")
		}
		return db
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic("failed to connect postgres database")
	}
	return db
}
