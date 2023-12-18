package db

import (
	"database/sql"
	"fmt"

	"github.com/antmordel/techtheon/foundation/config"
)

func Connect() (*sql.DB, error) {
	// Fetch the environment variables and check if they are set
	user, err := config.GetEnvVar("POSTGRES_USER")
	if err != nil {
		return nil, err
	}

	password, err := config.GetEnvVar("POSTGRES_PASSWORD")
	if err != nil {
		return nil, err
	}

	dbname, err := config.GetEnvVar("POSTGRES_DB")
	if err != nil {
		return nil, err
	}

	host, err := config.GetEnvVar("POSTGRES_HOST")
	if err != nil {
		return nil, err
	}

	port, err := config.GetEnvVar("POSTGRES_PORT")
	if err != nil {
		return nil, err
	}

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", user, password, dbname, host, port)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}
