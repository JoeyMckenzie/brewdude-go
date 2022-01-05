package repository

import "database/sql"

func (br breweryRepository) GetBreweries() ([]Brewery, error) {
	var breweries []Brewery

	const getBreweriesByIdSql = `
SELECT id, 
	   created_at, 
	   updated_at, 
	   image_url, 
	   name, 
	   street_address,
	   street_address_extended,
       city,
       state,
       zip_code,
       zip_code_extension
FROM public.breweries`

	rows, err := br.db.Query(getBreweriesByIdSql)
	if err != nil {
		return breweries, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			br.logger.ErrorWhile("closing the database connection", err)
		}
	}(rows)

	for rows.Next() {
		brewery := Brewery{}

		if err := rows.
			Scan(&brewery.Id,
				&brewery.CreatedAt,
				&brewery.UpdatedAt,
				&brewery.ImageUrl,
				&brewery.Name,
				&brewery.StreetAddress,
				&brewery.StreetAddressExtended,
				&brewery.City,
				&brewery.State,
				&brewery.ZipCode,
				&brewery.ZipCodeExtension);
		err != nil {
			return breweries, err
		}

		breweries = append(breweries, brewery)
	}

	return breweries, err
}
