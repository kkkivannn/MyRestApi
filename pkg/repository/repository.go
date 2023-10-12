package repository

import (
	api "RestaurantRestApi"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateClient(client api.Client) (int, error)
	GetClient(phone, password string) (api.Client, error)
	GetClientById(clientId int) (api.Client, error)
	GetProfile(clientId int) (api.Profile, error)
}

type AuthorizationEmployee interface {
	CreateEmployee(employee api.Employee) (int, error)
	GetEmployeeProfile(idEmployee int) (api.EmployeeProfile, error)
}

type Repository struct {
	Authorization
	AuthorizationEmployee
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:         NewAuthorizationRepository(db),
		AuthorizationEmployee: NewAuthorizationEmployeeRepository(db),
	}
}
