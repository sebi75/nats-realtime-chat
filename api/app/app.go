package app

import (
	"api/app/auth"
	"api/app/chat"
	"api/app/messageBroker"
	"api/app/ping"
	"api/env"
	"api/pkg/nats"
	"api/utils/logger"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Start and manage the application
func Start() {
	config := env.GetConfig()
	config.ConfigSanityCheck()
	natsClient, err := nats.New(config.NATS.Url)
	if err != nil {
		logger.Error("Error connecting to the NATS Client")
		panic(err)
	}
	router := mux.NewRouter()

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	messageBroker := messageBroker.New(natsClient)
	authService := auth.NewAuthService(config)
	authHandler := auth.NewAuthHandlers(authService)
	connectHandler, err := chat.NewChatHandler(messageBroker)
	if err != nil {
		log.Fatal(err)
	}

	router.HandleFunc("/ping", ping.Ping).Methods(http.MethodGet).Name("ping")
	router.HandleFunc("/connect", connectHandler.Connect).Methods(http.MethodGet).Name("connect")
	router.HandleFunc("/auth/login", authHandler.Signin).Methods(http.MethodPost).Name("login")
	router.HandleFunc("/auth/signup", authHandler.Signup).Methods(http.MethodPost).Name("signup")
	router.HandleFunc("/auth/verify", authHandler.Verify).Methods(http.MethodGet).Name("verify")

	http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router))
}
