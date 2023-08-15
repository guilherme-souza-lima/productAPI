package infra

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Environment        string
	App                string
	PostgresqlHost     string
	PostgresqlPort     string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDatabase string
}

func NewConfig() Config {
	if os.Getenv("ENVIRONMENT") == "" {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatalln("Error loading env file")
		}
	}

	return Config{
		Environment:        os.Getenv("APP"),
		App:                os.Getenv("ENVIRONMENT"),
		PostgresqlHost:     os.Getenv("HOST_POSTGRES"),
		PostgresqlPort:     os.Getenv("PORT_POSTGRES"),
		PostgresqlUser:     os.Getenv("USER_POSTGRES"),
		PostgresqlPassword: os.Getenv("PASSWORD_POSTGRES"),
		PostgresqlDatabase: os.Getenv("DATABASE_POSTGRES"),
	}
}
