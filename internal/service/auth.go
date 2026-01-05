package service

import (
	"crypto/sha256"
	"encoding/hex"
	"gachibank/domain"
	"gachibank/internal/repository"
)

type AuthService struct {
	rep repository.Authorization
}

func NewAuthService(rep repository.Authorization) *AuthService {
	return &AuthService{rep: rep}
}

func (s *AuthService) CreateUser(user domain.User) (int, error) {
	user.Password = generatePassword(user.Password)
	return s.rep.CreateUser(user)
}

func generatePassword(password string) string {
	hash := sha256.Sum256([]byte(password))

	hashString := hex.EncodeToString(hash[:])

	return hashString
}
