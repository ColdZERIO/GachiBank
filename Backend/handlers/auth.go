package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"gachibank/Backend/models"
	"io"
	"log"
	"net/http"
	"os"
)

func FrontAuthHandler(w http.ResponseWriter, r *http.Request) {
	front := `
<html>
<head>
    <title>Go Form</title>
</head>
<body>
    <form action="/auth/success" method="POST">

		<!-- Login -->
        <label for="loginField">Login:</label>
        <input 
            type="login"
            id="loginField" 
            name="login"
            required
        >
        <br><br>>

        <!-- Password -->
        <label for="passwordField">Password:</label>
        <input 
            type="password" 
            id="passwordField" 
            name="password" 
            required
        >
        <br><br>

        <button type="submit">Apply</button>
    </form>
</body>
</html>`

	fmt.Fprint(w, front)
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var (
		users      = models.DataBase{}
		foundLogin = false
		fileName   = "./database/database.json"
	)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Err open file")
		return
	}

	db, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Err read the file")
		return
	}

	err = json.Unmarshal(db, &users)
	if err != nil {
		log.Fatal("Err parse json file")
		return
	}

	login := r.FormValue("login")
	password := r.FormValue("password")

	for _, user := range users.Base {
		if user.Login == login {
			foundLogin = true

			hash := sha256.Sum256([]byte(password))
			hashStr := hex.EncodeToString(hash[:])
			if hashStr == user.Password {
				msg := fmt.Sprintf("Hello, %s!", user.Login)
				w.Write([]byte(msg))
				break
			}

			w.Write([]byte("Incorrect password\nTry again"))
		}
	}

	if !foundLogin {
		w.Write([]byte("User is not found"))
	}
}
