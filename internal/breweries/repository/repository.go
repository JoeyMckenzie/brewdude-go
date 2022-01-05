package repository

import (
	"database/sql"
	"github.com/joeymckenzie/brewdude/pkg/logging"
)

type BreweryRepository interface {
	GetBrewery(id string) (Brewery, error)
	GetBreweries() ([]Brewery, error)
	CreateBrewery(brewery *Brewery) error
}

type breweryRepository struct {
	db *sql.DB
	logger logging.Logger
}

func NewBreweryRepository(db *sql.DB, logger logging.Logger) BreweryRepository {
	return breweryRepository{db, logger}
}
