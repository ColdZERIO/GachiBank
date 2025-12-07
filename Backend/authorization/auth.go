package authorization

import (
	"encoding/json"
	"gachibank/database"
	"net/http"
)

type UserAuth struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// Вынес как глобальную переменную для будующего обращения и для иницилизации в мейне (а так по-любому нужны структуры)
var db *database.Database

// Иницилизация БД для мейна
func Init(dbInstance *database.Database) {
	db = dbInstance
}

func UserRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != http.MethodPost {
		http.Error(w, "incorrect method", http.StatusMethodNotAllowed)
		return
	}

	var User UserAuth

	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		http.Error(w, "pars failed", http.StatusBadRequest)
		return
	}

	if len(User.Login) < 2 && User.Login == "" {
		http.Error(w, "invalid user name", http.StatusBadRequest)
		return
	}

	if len(User.Password) < 6 && User.Password == "" {
		http.Error(w, "invalid password", http.StatusBadRequest)
		return
	}

	dataDB, err := db.GetUserAuth(User.Login, User.Password)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	if dataDB == nil {
		http.Error(w, "user not found in DB", http.StatusNotFound)
		return
	}

	reqDB := map[string]string{
		"login":    User.Login,
		"password": User.Password,
	}

	json.NewEncoder(w).Encode(reqDB)
}

func GetAuth(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}
