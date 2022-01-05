package shared

import (
	"database/sql"
	"github.com/joeymckenzie/brewdude/internal/breweries/domain"
	"github.com/joeymckenzie/brewdude/pkg/logging"
)

type ServiceInterface interface {
	New(logger logging.Logger, db *sql.DB) BreweriesService
}

type BreweriesService interface {
	ServiceInterface
	CreateBrewery(request domain.BreweryDto) (domain.BreweryDto, error)
	GetBrewery(id string) (domain.BreweryDto, error)
}

type BeersService interface {
	ServiceInterface
	CreateBeer() error
	GetBeer(id string) error
}
