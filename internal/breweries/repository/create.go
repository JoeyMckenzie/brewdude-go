package repository

func (br breweryRepository) CreateBrewery(brewery *Brewery) error {
	const insertBrewerySql = `
INSERT INTO public.breweries (id, created_at, updated_at, image_url, name, street_address, street_address_extended, city, state, zip_code, zip_code_extension)
VALUES (gen_random_uuid(), current_timestamp, current_timestamp, $1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, created_at, updated_at;
`

	err := br.db.
		QueryRow(insertBrewerySql,
			brewery.ImageUrl,
			brewery.Name,
			brewery.StreetAddress,
			brewery.StreetAddressExtended,
			brewery.City,
			brewery.State,
			brewery.ZipCode,
			brewery.ZipCodeExtension).
		Scan(&brewery.Id,
			&brewery.CreatedAt,
			&brewery.UpdatedAt)

	return err
}
