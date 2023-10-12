package service

import (
	api "RestaurantRestApi"
	"RestaurantRestApi/pkg/auth"
	"RestaurantRestApi/pkg/repository"
	"crypto/sha1"
	"fmt"
)

type Authorization interface {
	CreateUser(clients api.Client) (int, error)
	GenerateTokens(phone, password string) (string, error)
	ParseToken(accessesToken string) (int, error)
	RefreshToken(refreshToken string) (string, error)
	GetClientById(clientId int) (api.Client, error)
	GetProfileClient(clientId int) (api.Profile, error)
}
type AuthorizationEmployee interface {
	CreateEmployee(employee api.Employee) (int, error)
	GetProfileEmployee(employeeId int) (api.EmployeeProfile, error)
}

type Service struct {
	Authorization
	AuthorizationEmployee
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization:         NewAuthorizationService(repository),
		AuthorizationEmployee: NewAuthorizationEmployeeService(repository),
	}
}

func PasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(auth.Salt)))
}
