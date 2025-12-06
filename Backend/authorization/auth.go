package authorization

import "net/http"

type UserAuth struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func UserRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "incorrect method", http.StatusMethodNotAllowed)
		return
	}

}
