package repo

import (
	"database/sql"
	"errors"
	"os"
)

func NewSQLDB(connectionString string, provider LoggerProvider) *sql.DB {
	logger := provider(nil)
	logger.Info("NewSQLDB", "connection string", connectionString)
	dir, err := os.Getwd()
	if err != nil {
		panic(errors.Join(err, errors.New("failed to get current working directory in NewSQLDB")))
	}
	logger.Info("NewSQLDB", "PWD", dir)
	db, err := sql.Open("sqlite3", connectionString)
	if err != nil {
		panic(errors.Join(err, errors.New("failed to open DB in NewSQLDB")))
	}
	return db
}
