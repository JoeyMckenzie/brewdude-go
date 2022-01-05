package mapper

import (
	"github.com/joeymckenzie/brewdude/internal/breweries/domain"
	"github.com/joeymckenzie/brewdude/internal/breweries/repository"
)

func MapEntityToDto(entity repository.Brewery) domain.BreweryDto {
	return domain.BreweryDto{
		Id:                    entity.Id,
		Name:                  entity.Name,
		StreetAddress:         entity.StreetAddress,
		StreetAddressExtended: entity.StreetAddressExtended,
		City:                  entity.City,
		State:                 entity.State,
		ZipCode:               entity.ZipCode,
		ZipCodeExtension:      entity.ZipCodeExtension,
	}
}
