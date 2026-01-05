package repository

import (
	"database/sql"
	"log"

	"github.com/TifaLuv/GolangServer/domain"
)

type AuthSQL struct {
	db *sql.DB
}

func NewAuthSQL(db *sql.DB) *AuthSQL {
	return &AuthSQL{db: db}
}

func (r *AuthSQL) CreateUser(user domain.User) (int, error) {
	res, err := r.db.Exec("INSERT INTO users (name, username, password) VALUES (:name, :username, :password)",
		sql.Named("name", user.Name),
		sql.Named("username", user.Username),
		sql.Named("password", user.Password))
	if err != nil {
		log.Printf("error create user")
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Printf("error get last ID")
	}
	return int(id), nil
}
