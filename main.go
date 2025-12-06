package main

import (
	"gachibank/Backend/authorization"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	router.Get("/auth", authorization.GetAuth)
	router.Post("/auth", authorization.UserRegister)

	router.Handle("/*", http.FileServer(http.Dir("./")))
}
