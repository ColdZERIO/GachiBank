package main

import (
	"gachibank/Backend/handlers"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
		return
	}

	router := chi.NewRouter()

	router.Post("/auth", handlers.AuthHandler)
	router.Post("/reg", handlers.RegistrationHandler)

	router.Handle("/*", http.FileServer(http.Dir("./")))

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
		return
	}
}
