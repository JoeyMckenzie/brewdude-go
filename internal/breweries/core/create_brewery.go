package core

import (
	"errors"
	"fmt"
	"github.com/joeymckenzie/brewdude/internal/breweries/domain"
	"github.com/joeymckenzie/brewdude/internal/breweries/mapper"
	"github.com/joeymckenzie/brewdude/internal/breweries/repository"
	"github.com/joeymckenzie/brewdude/internal/breweries/validation"
)

func (bs breweryService) CreateBrewery(request domain.BreweryDto) (domain.BreweryDto, error) {
	var breweryDto domain.BreweryDto

	// 2. ValidateBreweryRequest fields
	if validationErrors := validation.ValidateBreweryRequest(request); len(validationErrors) > 0 {
		return breweryDto, errors.New(validationErrors[0])
	}

	// 3. Persist brewery with addressRepository
	brewery := repository.Brewery{
		Name:                  request.Name,
		ImageUrl:              request.ImageUrl,
		StreetAddress:         request.StreetAddress,
		StreetAddressExtended: request.StreetAddressExtended,
		City:                  request.City,
		State:                 request.State,
		ZipCode:               request.ZipCode,
		ZipCodeExtension:      request.ZipCodeExtension,
	}

	err := bs.repository.CreateBrewery(&brewery)

	if err != nil {
		bs.Logger.ErrorWhile(fmt.Sprintf("creating brewery %s", request.Name), err)
		return breweryDto, err
	}

	return mapper.MapEntityToDto(brewery), nil


}
