package repository

import (
	api "RestaurantRestApi"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AuthorizationRepository struct {
	db *sqlx.DB
}

func NewAuthorizationRepository(db *sqlx.DB) *AuthorizationRepository {
	return &AuthorizationRepository{db: db}
}

func (a *AuthorizationRepository) CreateClient(client api.Client) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, age, phone_number, password_hash) VALUES ($1, $2, $3, $4) RETURNING id", clientsTable)
	row := a.db.QueryRow(query, client.Name, client.Age, client.Phone, client.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (a *AuthorizationRepository) GetClient(phone, password string) (api.Client, error) {
	var client api.Client
	query := fmt.Sprintf("SELECT id FROM %s WHERE phone_number=$1 AND password_hash=$2", clientsTable)
	if err := a.db.Get(&client, query, phone, password); err != nil {
		return client, err
	}
	return client, nil
}

func (a *AuthorizationRepository) GetClientById(clientId int) (api.Client, error) {
	var client api.Client
	query := fmt.Sprintf("SELECT tl.id, tl.name, tl.age, tl.phone_number FROM %s tl WHERE tl.id = $1", clientsTable)
	if err := a.db.Get(&client, query, clientId); err != nil {
		return api.Client{}, err
	}
	return client, nil
}

func (a *AuthorizationRepository) GetProfile(clientId int) (api.Profile, error) {
	var client api.Profile
	query := fmt.Sprintf("SELECT tl.name, tl.age, tl.phone_number FROM %s tl WHERE tl.id = $1", clientsTable)
	if err := a.db.Get(&client, query, clientId); err != nil {
		return api.Profile{}, err
	}
	return client, nil
}
