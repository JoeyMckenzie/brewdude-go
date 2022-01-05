package api

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/joeymckenzie/brewdude/internal/breweries/core"
	"github.com/joeymckenzie/brewdude/pkg/logging"
)

type BreweriesController struct {
	Router  *mux.Router
	Logger  logging.Logger
	Db      *sql.DB
	service core.BreweryService
}

func NewBreweriesController(router *mux.Router, logger logging.Logger, db *sql.DB) BreweriesController {
	return BreweriesController{
		router, logger, db, core.NewBreweryService(logger, db),
	}
}
