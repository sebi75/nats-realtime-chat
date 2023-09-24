package main

import (
	"auth-service/app"
	"auth-service/utils/logger"
)

func main() {
	logger.Info("Starting auth-service")
	app.Start()
}
