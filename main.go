package main

import (
	"gachibank/Backend/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	router := chi.NewRouter()

	router.Post("/auth", handlers.AuthHandler)
	router.Post("/reg", handlers.RegistrationHandler)

	router.Handle("/*", http.FileServer(http.Dir("./")))

	http.ListenAndServe(":8080", router)
}
