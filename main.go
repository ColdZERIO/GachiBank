package main

import (
	"context"
	"gachibank/Backend/authorization"
	"gachibank/database"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db, err := database.DatabaseInit()
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
	defer db.Client.Disconnect(context.Background())

	router := chi.NewRouter()

	router.Get("/auth", authorization.GetAuth)
	router.Post("/auth", authorization.UserRegister)

	router.Handle("/*", http.FileServer(http.Dir("./")))

	http.ListenAndServe(":8080", router)
}
