package validation

import (
	"github.com/joeymckenzie/brewdude/internal/breweries/domain"
	"github.com/joeymckenzie/brewdude/internal/shared"
	"regexp"
)

func ValidateBreweryRequest(request domain.BreweryDto) []string {
	var validationErrors []string

	if request.Name == "" {
		validationErrors = append(validationErrors, "Please provide a brewery name.")
	}

	if request.StreetAddress == "" {
		validationErrors = append(validationErrors, "Please provide a street address.")
	}

	if request.City == "" {
		validationErrors = append(validationErrors, "Please provide a city.")
	}

	if request.State == "" {
		validationErrors = append(validationErrors, "Please provide a state code.")
	}

	validMatch, err := regexp.Match(shared.ValidStateAbbreviationRegex, []byte(request.State))

	if !validMatch || err != nil {
		validationErrors = append(validationErrors, "Please provide a valid state code.")
	}

	if request.ZipCode == "" {
		validationErrors = append(validationErrors, "Please provide a zip code.")
	}

	return validationErrors
}
