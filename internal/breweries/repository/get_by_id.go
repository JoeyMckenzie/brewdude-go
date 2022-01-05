package repository

func (br breweryRepository) GetBrewery(id string) (Brewery, error) {
	var brewery Brewery

	const getBreweryByIdSql = `
SELECT *
FROM public.breweries b
WHERE b.id = $1`

	err := br.db.
		QueryRow(getBreweryByIdSql, id).
		Scan(&brewery.Id,
			&brewery.CreatedAt,
			&brewery.UpdatedAt,
			&brewery.ImageUrl,
			&brewery.Name,
			&brewery.Id,
			&brewery.CreatedAt,
			&brewery.UpdatedAt,
			&brewery.StreetAddress,
			&brewery.StreetAddressExtended,
			&brewery.City,
			&brewery.State,
			&brewery.ZipCode,
			&brewery.StreetAddressExtended)

	if err != nil {
		return brewery, err
	}

	return brewery, nil
}
