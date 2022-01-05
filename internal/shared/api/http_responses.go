package api

import (
	"encoding/json"
	shared "github.com/joeymckenzie/brewdude/internal/shared/domain"
	"github.com/joeymckenzie/brewdude/pkg/logging"
	"net/http"
)

func writeErrorResponse(w http.ResponseWriter, err error, logger logging.Logger) {
	marshalledErrorResponse, err := json.Marshal(shared.BrewdudeApiResponse{
		Success: false,
		Errors:  []string{err.Error()},
	})

	if err != nil {
		logger.FatalWhile("marshalling the error response", err)
	}

	w.Write(marshalledErrorResponse)
}

func UnsupportedMediaType(w http.ResponseWriter, err error, logger logging.Logger) {
	w.WriteHeader(http.StatusUnsupportedMediaType)
	writeErrorResponse(w, err, logger)
}

func BadRequest(w http.ResponseWriter, err error, logger logging.Logger) {
	w.WriteHeader(http.StatusBadRequest)
	writeErrorResponse(w, err, logger)
}

func InternalServerError(w http.ResponseWriter, err error, logger logging.Logger) {
	w.WriteHeader(http.StatusInternalServerError)
	writeErrorResponse(w, err, logger)
}
