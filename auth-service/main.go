package main

import (
	domain "auth-service/domain/auth"
	"auth-service/env"
	"auth-service/lib"
	"log"
)

func main() {
	config := env.GetConfig()
	config.ConfigSanityCheck()

	db, err := lib.InitDbClient(config)
	if err != nil {
		log.Fatal(err)
	}
	_ = domain.NewAuthRepositoryDB(db)
}
