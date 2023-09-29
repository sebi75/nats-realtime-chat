package app

import (
	hs "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	// import app modules
	auth "auth-service/app/auth"
	"auth-service/env"
	"auth-service/lib"
	"log"
	"net/http"
)

func Start() {
	// implement all the code that starts and orchestrates the application
	config := env.GetConfig()
	config.ConfigSanityCheck()

	headersOk := hs.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := hs.AllowedOrigins([]string{"*"})
	methodsOk := hs.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	db, err := lib.InitDbClient(config)
	if err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	authRepository := auth.NewAuthRepositoryDB(db)
	authService := auth.NewAuthService(authRepository)
	authHandlers := auth.NewAuthHandlers(authService)

	r.HandleFunc("/signup", authHandlers.Signup).Methods(http.MethodPost).Name("Signup")
	r.HandleFunc("/signin", authHandlers.Signin).Methods(http.MethodPost).Name("Signin")
	r.HandleFunc("/verify", authHandlers.Verify).Methods(http.MethodGet).Name("Verify")

	http.ListenAndServe(":8083", hs.CORS(originsOk, headersOk, methodsOk)(r))
}
