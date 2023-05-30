package env

import (
	"api/utils/logger"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	logger.Info("Loading environment variables...")
	godotenv.Load()
	logger.Info("Environment variables loaded.")
}

type NATSConfig struct {
	Url string
}

type DBConfig struct {
	Host         string
	Port         string
	User         string
	Password     string
	DatabaseName string
}

type Config struct {
	NATS NATSConfig
	DB   DBConfig
}

func GetConfig() *Config {
	return &Config{
		NATS: NATSConfig{
			Url: os.Getenv("NATS_URL"),
		},
		DB: DBConfig{
			Host:         os.Getenv("DB_HOST"),
			Port:         os.Getenv("DB_PORT"),
			User:         os.Getenv("DB_USER"),
			Password:     os.Getenv("DB_PASSWORD"),
			DatabaseName: os.Getenv("DB_NAME"),
		},
	}
}

func (c *Config) ConfigSanityCheck() {
	if c.NATS.Url == "" {
		log.Fatal("NATS_URL is not set")
	}

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

	logger.Info("Environment variables sanity check passed.")
}
