package core

import (
	"database/sql"
	"github.com/joeymckenzie/brewdude/internal/breweries/repository"
	"github.com/joeymckenzie/brewdude/internal/shared"
	"github.com/joeymckenzie/brewdude/pkg/logging"
)

type breweryService struct {
	Logger     logging.Logger
	repository repository.BreweryRepository
}

func (bs breweryService) New(logger logging.Logger, db *sql.DB) shared.BreweriesService {
	return breweryService{logger, repository.NewBreweryRepository(db, logger)}
}
