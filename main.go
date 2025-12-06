package main

import (
	"gachibank/Backend/authorization"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	router.Post("/auth", authorization.UserRegister)
}