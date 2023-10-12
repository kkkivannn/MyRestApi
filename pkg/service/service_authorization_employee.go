package service

import (
	api "RestaurantRestApi"
	"RestaurantRestApi/pkg/repository"
)

type AuthorizationEmployeeService struct {
	repository *repository.Repository
}

func NewAuthorizationEmployeeService(repository *repository.Repository) *AuthorizationEmployeeService {
	return &AuthorizationEmployeeService{repository: repository}
}

func (a *AuthorizationEmployeeService) CreateEmployee(employee api.Employee) (int, error) {
	employee.PasswordHash = PasswordHash(employee.PasswordHash)
	return a.repository.CreateEmployee(employee)
}

func (a *AuthorizationEmployeeService) GetProfileEmployee(employeeId int) (api.EmployeeProfile, error) {
	return a.repository.GetEmployeeProfile(employeeId)
}
