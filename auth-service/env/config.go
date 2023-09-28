package env

import (
	"auth-service/utils/logger"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	logger.Info("Loading environment variables...")
	godotenv.Load()
	logger.Info("Environment variables loaded.")
}

type DBConfig struct {
	Host         string
	Port         string
	User         string
	Password     string
	DatabaseName string
}

type RestConfig struct {
	Jwt_Secret string
}

type Config struct {
	DB   DBConfig
	Rest RestConfig
}

func GetConfig() *Config {
	return &Config{
		DB: DBConfig{
			Host:         os.Getenv("DB_HOST"),
			Port:         os.Getenv("DB_PORT"),
			User:         os.Getenv("DB_USER"),
			Password:     os.Getenv("DB_PASSWORD"),
			DatabaseName: os.Getenv("DB_NAME"),
		},
		Rest: RestConfig{
			Jwt_Secret: os.Getenv("JWT_SECRET"),
		},
	}
}

func (c *Config) ConfigSanityCheck() {
	if c.DB.Host == "" {
		log.Fatal("DB_HOST is not set")
	}

	if c.DB.Port == "" {
		log.Fatal("DB_PORT is not set")
	}

	if c.DB.User == "" {
		log.Fatal("DB_USER is not set")
	}

	if c.DB.Password == "" {
		log.Fatal("DB_PASSWORD is not set")
	}

	if c.DB.DatabaseName == "" {
		log.Fatal("DB_NAME is not set")
	}
	if c.Rest.Jwt_Secret == "" {
		log.Fatal("JWT_SECRET is not set")
	}

	logger.Info("Environment variables sanity check passed.")
}
