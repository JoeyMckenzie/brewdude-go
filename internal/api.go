package internal

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/joeymckenzie/brewdude/internal/breweries/api"
	"github.com/joeymckenzie/brewdude/pkg/logging"
)

func BuildApiRouter(logger logging.Logger, db *sql.DB) *mux.Router {
	router := mux.NewRouter().
		PathPrefix("/api/v1").
		Subrouter()

	breweriesController := api.NewBreweriesController(router.PathPrefix("/breweries").Subrouter(), logger, db)
	breweriesController.RegisterRoutes()

	return router
}