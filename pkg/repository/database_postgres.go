package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	clientsTable  = "clients"
	employeeTable = "employees"
	roleTable     = "Roles"
)

type ConfigDb struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
	SSLMode  string
}

func NewPostgresDb(config *ConfigDb) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", config.Host, config.Port, config.Username, config.DbName, config.Password, config.SSLMode))
	if err != nil {
		return nil, err
	}
	return db, nil
}
