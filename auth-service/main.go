package main

import (
	domain "auth-service/domain/auth"
	"auth-service/env"
	"auth-service/lib"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	config := env.GetConfig()
	config.ConfigSanityCheck()

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	db, err := lib.InitDbClient(config)
	if err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	_ = domain.NewAuthRepositoryDB(db)

	http.ListenAndServe(":8083", handlers.CORS(originsOk, headersOk, methodsOk)(r))
}
