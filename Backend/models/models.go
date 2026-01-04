package models

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// База даннех пользователей
type DataBase struct {
	Base []User `json:"base"`
}

// Параметры пользователя
type User struct {
	Login    string `json:"login"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	BirthDay string `json:"birthday"`
	Age      int    `json:"age"`
	Subject  string `json:"subject"`
	Date     string `json:"date"`
}

// Хеширование пароля
func (u *User) SetPasswordHash(key string) string {
	hash := sha256.Sum256([]byte(key))
	return hex.EncodeToString(hash[:])
}

// Определение возраста по дате рождения
func (u *User) SetAge(date string) int {
	age, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Println("invalid date format", age, "\n", date)
		return 0
	}

	ageInt := time.Since(age)
	return int(ageInt.Hours() / 24 / 365)
}

// Создание jwt токена
func (u *User) SetJwtToken(secretKey string) (string, error) {
	claims := jwt.MapClaims{
		"login":    u.Login,
		"password": u.Password,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

// Сохранение пользователя в файл json
func (u *DataBase) SaveFile(file string) error {
	data, err := json.MarshalIndent(u, "", "	")
	if err != nil {
		return err
	}

	return os.WriteFile(file, data, 0755)
}

// Добавление пользователя в слайс базы данных и обновление базы данных
func (u *DataBase) AddUser(newUser User, filename string) error {
	for _, user := range u.Base {
		if user.Login == newUser.Login {
			return fmt.Errorf("user with Login %s already exists", newUser.Login)
		}
	}

	u.Base = append(u.Base, newUser)

	return u.SaveFile(filename)
}

// Выгрузка данных из базы данных
func (db *DataBase) LoadFile(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			db.Base = []User{}
			return nil
		}
		return err
	}
	if len(file) == 0 {
		db.Base = []User{}
		return nil
	}
	return json.Unmarshal(file, db)
}

// Обновляем слайс пользовтелей новыми данными
func NewUsersDB(filename string) (*DataBase, error) {
	database := &DataBase{}
	err := database.LoadFile(filename)
	return database, err
}
