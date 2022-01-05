package domain

type BreweryDto struct {
	Id                    string `json:"id"`
	Name                  string `json:"name"`
	ImageUrl              string `json:"imageUrl,omitempty"`
	StreetAddress         string `json:"streetAddress"`
	StreetAddressExtended string `json:"StreetAddressExtended,omitempty"`
	City                  string `json:"city"`
	State                 string `json:"state"`
	ZipCode               string `json:"zipCode"`
	ZipCodeExtension      string `json:"zipCodeExtension,omitempty"`
}
