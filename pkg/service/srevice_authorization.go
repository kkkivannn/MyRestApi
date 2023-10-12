package service

import (
	api "RestaurantRestApi"
	"RestaurantRestApi/pkg/auth"
	"RestaurantRestApi/pkg/repository"
)

type AuthorizationService struct {
	repository *repository.Repository
	manager    *auth.Manager
}

func NewAuthorizationService(repository *repository.Repository) *AuthorizationService {
	return &AuthorizationService{repository: repository}
}

func (a *AuthorizationService) CreateUser(client api.Client) (int, error) {
	client.Password = PasswordHash(client.Password)
	return a.repository.CreateClient(client)
}

func (a *AuthorizationService) GenerateTokens(phone, password string) (string, error) {
	client, err := a.repository.GetClient(phone, PasswordHash(password))
	if err != nil {
		return "", err
	}
	token, err := a.manager.NewTokenGenerate(client.Id)
	if err != nil {
		return "", err
	}

	return token, nil

}
func (a *AuthorizationService) ParseToken(accessToken string) (int, error) {
	id, err := a.manager.ParseToken(accessToken)
	if err != nil {
		return 0, err
	}
	return id, err
}

func (a *AuthorizationService) RefreshToken(token string) (string, error) {
	id, err := a.manager.ParseToken(token)
	if err != nil {
		return "", err
	}
	client, err := a.repository.GetClientById(id)
	if err != nil {
		return "", err
	}
	NewToken, err := a.manager.NewTokenGenerate(client.Id)
	if err != nil {
		return "", err
	}
	return NewToken, nil
}
func (a *AuthorizationService) GetClientById(clientId int) (api.Client, error) {
	return a.repository.GetClientById(clientId)
}

func (a *AuthorizationService) GetProfileClient(clientId int) (api.Profile, error) {
	return a.repository.GetProfile(clientId)
}
