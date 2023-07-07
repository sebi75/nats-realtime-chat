package lib

import (
	"auth-service/env"
	"auth-service/utils/logger"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func InitDbClient(config *env.Config) (*sqlx.DB, error) {
	connectionString := config.DB.User + ":" + config.DB.Password + "@tcp(" + config.DB.Host + ":" + config.DB.Port + ")/" + config.DB.DatabaseName
	dbClient, err := sqlx.Connect("mysql", connectionString)

	dbClient.SetConnMaxLifetime(time.Minute * 1)
	if err != nil {
		logger.Error("InitDbClient::Failed to connect to database.")
		return nil, err
	}

	logger.Info("InitDbClient::Successfully connected to database.")

	return dbClient, nil
}
