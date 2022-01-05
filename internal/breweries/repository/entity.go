package repository

import (
	"github.com/joeymckenzie/brewdude/pkg/persistence"
)

type Brewery struct {
	persistence.BaseAuditableEntity
	Name                  string
	ImageUrl              string
	StreetAddress         string
	StreetAddressExtended string
	City                  string
	State                 string
	ZipCode               string
	ZipCodeExtension      string
}
