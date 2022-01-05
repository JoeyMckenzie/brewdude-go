package api

import (
	"fmt"
	"github.com/joeymckenzie/brewdude/internal/shared"
	"net/http"
)

var idPathMatcher = fmt.Sprintf("/{id:%s}", shared.UuidRegex)

func (bc *BreweriesController) RegisterRoutes() {
	bc.Logger.Info("Registering breweries API routes...")

	// GET /breweries
	// bc.Router.HandleFunc("", bc.GetBreweries).
	// 	Methods(http.MethodGet)
	// bc.Logger.Info("Registered route GET /breweries")

	// GET /breweries/<id>
	bc.Router.HandleFunc(idPathMatcher, bc.GetBrewery).
		Methods(http.MethodGet)
	bc.Logger.Info("Registered route GET /breweries/<id>")

	// POST /breweries
	bc.Router.HandleFunc("", bc.CreateBrewery).
		Methods(http.MethodPost)
	bc.Logger.Info("Registered route POST /breweries")

	bc.Logger.Info("Brewery routes registered!")
}
