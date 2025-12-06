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

// type Auth struct {
// 	database *database.Database
// }

// func (a Auth) CheckUserAuth(login, password string) (models.UserAuthDB, error) {
// 	userCheck, err := a.database.GetUserAuth(login, password)
// 	if err != nil {
// 		return models.UserAuthDB, err
// 	}
// }

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
	}

	if len(User.Login) < 2 && User.Login == "" {
		http.Error(w, "invalid user name", http.StatusBadRequest)
	}

	if len(User.Password) < 6 && User.Password == "" {
		http.Error(w, "invalid password", http.StatusBadRequest)
	}

	var data *database.Database

	dataDB, err := data.GetUserAuth(User.Login, User.Password)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
	}

	if dataDB == nil {
		http.Error(w, "user not found in DB", http.StatusNotFound)
	}

	reqDB := map[string]string{
		"login":    User.Login,
		"password": User.Password,
	}

	json.NewEncoder(w).Encode(reqDB)
}

func GetAuth(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../index.html")
}
