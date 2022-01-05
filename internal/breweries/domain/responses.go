package domain

type GetBreweryResponse struct {
	Brewery BreweryDto `json:"brewery"`
}

type GetBreweriesResponse struct {
	Breweries []BreweryDto `json:"breweries"`
}

type BreweryResponse struct {
	Brewery BreweryDto `json:"brewery"`
}