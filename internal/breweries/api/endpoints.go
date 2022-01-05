package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joeymckenzie/brewdude/internal/breweries/domain"
	"github.com/joeymckenzie/brewdude/internal/shared/api"
	shared "github.com/joeymckenzie/brewdude/internal/shared/domain"
	"github.com/joeymckenzie/brewdude/pkg/logging"
	"net/http"
)

func handleServiceResponse(w http.ResponseWriter, status int, brewery domain.BreweryDto, logger logging.Logger) {
	marshalledResponse, err := json.Marshal(shared.BrewdudeApiResponse{
		Success: true,
		Data:    domain.BreweryResponse{Brewery: brewery},
	})

	if err != nil {
		api.InternalServerError(w, err, logger)
		return
	}

	w.WriteHeader(status)
	w.Write(marshalledResponse)
}

func (bc *BreweriesController) CreateBrewery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var createBreweryRequest domain.CreateBreweryRequest

	err := json.
		NewDecoder(r.Body).
		Decode(&createBreweryRequest)

	if err != nil {
		api.UnsupportedMediaType(w, err, bc.Logger)
		return
	}

	brewery, err := bc.service.CreateBrewery(createBreweryRequest.Brewery)

	if err != nil {
		api.BadRequest(w, err, bc.Logger)
		return
	}

	handleServiceResponse(w, http.StatusCreated, brewery, bc.Logger)
}

func (bc *BreweriesController) GetBrewery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	breweryId := mux.Vars(r)["id"]

	bc.Logger.Info(fmt.Sprintf("Request received to retrieve brewery %s", breweryId))

	brewery, err := bc.service.GetBrewery(breweryId)

	if err != nil {
		api.BadRequest(w, err, bc.Logger)
		return
	}

	handleServiceResponse(w, http.StatusOK, brewery, bc.Logger)
}
