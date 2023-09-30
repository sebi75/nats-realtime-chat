package app

import (
	"api/app/connect"
	"api/app/ping"
	"api/env"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Start and manage the application
func Start() {
	config := env.GetConfig()
	config.ConfigSanityCheck()
	// _, err := nats.New(cfg.NATS.Url)
	// if err != nil {
	// panic(err)
	// }
	router := mux.NewRouter()

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	connectHandler, err := connect.NewConnectHandler()
	if err != nil {
		log.Fatal(err)
	}

	router.HandleFunc("/ping", ping.Ping).Methods(http.MethodGet)
	router.HandleFunc("/connect", connectHandler.Connect).Methods(http.MethodGet)

	http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router))
}
