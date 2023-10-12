package repository

import (
	api "RestaurantRestApi"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AuthorizationEmployeeRepository struct {
	db *sqlx.DB
}

func NewAuthorizationEmployeeRepository(db *sqlx.DB) *AuthorizationEmployeeRepository {
	return &AuthorizationEmployeeRepository{db: db}
}

func (a *AuthorizationEmployeeRepository) CreateEmployee(employee api.Employee) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (login, name, surname, middle_name, password_hash) VALUES ($1, $2, $3, $4, $5) RETURNING id_employee", employeeTable)
	row := a.db.QueryRow(query, employee.Login, employee.Name, employee.Surname, employee.MiddleName, employee.PasswordHash)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (a *AuthorizationEmployeeRepository) GetEmployeeProfile(employeeId int) (api.EmployeeProfile, error) {
	var employee api.EmployeeProfile
	query := fmt.Sprintf("SELECT tl.name, tl.login, tl.surname, tl.middlename FROM %s tl WHERE tl.id_employee = $1", employeeTable)
	if err := a.db.Get(&employee, query, employeeId); err != nil {
		return api.EmployeeProfile{}, err
	}
	return employee, nil
}
