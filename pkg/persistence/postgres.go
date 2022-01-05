package persistence

import (
	"database/sql"
	"github.com/joeymckenzie/brewdude/pkg/logging"
	_ "github.com/lib/pq"
)

func InitializePostgres(logger logging.Logger, connectionString string) (*sql.DB, error) {
	logger.Info("Initializing postgres...")
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		logger.FatalWhile("An error occurred while initializing postgres.", err)
		return nil, err
	}

	logger.Info("Database connection established! Pinging server...")

	if err := db.Ping(); err != nil {
		logger.FatalWhile("verifying connection to database", err)
		return nil, err
	}

	logger.Info("Postgres is ready!")

	return db, nil
}
