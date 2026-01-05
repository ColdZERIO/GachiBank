package main

import (
	"gachibank/Backend/handlers"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
		return
	}

	http.HandleFunc("/reg", handlers.FrontRegHandler)
	http.HandleFunc("/reg/success", handlers.RegistrationHandler)
	http.HandleFunc("/auth", handlers.FrontAuthHandler)
	http.HandleFunc("/auth/success", handlers.AuthHandler)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}
