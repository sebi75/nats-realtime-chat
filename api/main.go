package main

import (
	"api/app"
	"api/utils/logger"
)

func main() {
	logger.Info("Starting API")
	app.Start()
}
