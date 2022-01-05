package core

import (
	"github.com/joeymckenzie/brewdude/internal/breweries/domain"
	"github.com/joeymckenzie/brewdude/internal/breweries/mapper"
)

func (bs breweryService) GetBrewery(id string) (domain.BreweryDto, error) {
	var mappedBrewery domain.BreweryDto

	brewery, err := bs.repository.GetBrewery(id)

	if err != nil {
		return mappedBrewery, err
	}

	mappedBrewery = mapper.MapEntityToDto(brewery)

	return mappedBrewery, nil
}
