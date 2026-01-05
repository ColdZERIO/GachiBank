package handlers

import (
	"fmt"
	"gachibank/internal/models"
	"net/http"
	"strings"
	"time"
)

func FrontRegHandler(w http.ResponseWriter, r *http.Request) {
	front := `
<html>
<head>
    <title>Go Form</title>
</head>
<body>
    <form action="/reg/success" method="POST">

        <!-- Name -->
        <label for="nameField">Name:</label>
        <input 
            type="name" 
            id="nameField" 
            name="name"
            minlength="2"
            required
        >
        <br><br>

		<!-- Login -->
        <label for="loginField">Login:</label>
        <input 
            type="login" 
            id="loginField" 
            name="login"
            minlength="6"
            required
        >
        <br><br>

        <!-- Subject -->
        <label for="subjectField">Subject:</label>
        <input 
            type="text" 
            id="textField" 
            name="subject"
            minlength="2"
            required
        >
        <br><br>

        <!-- Email -->
        <label for="emailField">Email:</label>
        <input 
            type="email" 
            id="emailField" 
            name="email" 
            minlength="6"
            required
        >
        <br><br>

        <!-- Password -->
        <label for="passwordField">Password:</label>
        <input 
            type="password" 
            id="passwordField" 
            name="password" 
            minlength="6"
            required
        >
        <br><br>

		<!-- BirthDay -->
        <label for="BirthDayField">BirthDayField:</label>
        <input 
            type="date" 
            id="BirthDayField" 
            name="birthday"
            required
        >
        <script> const today = new Date().toISOString().split('T') </script>
        <br><br>

        <button type="submit">Apply</button>
    </form>
</body>
</html>`

	fmt.Fprint(w, front)
}

// Регистрация пользователя
func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	var fileName = "./database/database.json"

	if r.Method != http.MethodPost {
		http.Error(w, "invalid method request", http.StatusMethodNotAllowed)
		return
	}
	// Вносим данные в поля
	user.Login = strings.ToLower(strings.Trim(r.FormValue("login"), " "))
	user.Name = strings.Trim(r.FormValue("name"), " ")
	user.Email = r.FormValue("email")
	user.Password = user.SetPasswordHash(r.FormValue("password"))
	user.BirthDay = r.FormValue("birthday")
	user.Age = user.SetAge(user.BirthDay)
	user.Subject = r.FormValue("subject")
	user.Date = time.Now().Format("2006-01-02 15:04")

	db, err := models.NewUsersDB(fileName)
	if err != nil {
		fmt.Printf("Err loading database: %v\n", err)
		return
	}

	err = db.AddUser(user, fileName)
	if err != nil {
		fmt.Println("Err to add new User")
		return
	}

	w.Write([]byte("Registration was successful!"))
}
