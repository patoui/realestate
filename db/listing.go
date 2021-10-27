package db

import (
	"database/sql"

	"github.com/patoui/realestate/models"
)

func (db Database) GetAllListings() (*models.ListingList, error) {
	list := &models.ListingList{}

	rows, err := db.Conn.Query("SELECT * FROM listings ORDER BY ID DESC")
	if err != nil {
		return list, err
	}

	for rows.Next() {
		var listing models.Listing
		err := rows.Scan(
			&listing.ID,
			&listing.MLSNumber,
			&listing.Address,
			&listing.Address2,
			&listing.City,
			&listing.PostalCode,
			&listing.State,
			&listing.Country,
			&listing.CreatedAt,
		)
		if err != nil {
			return list, err
		}
		list.Listings = append(list.Listings, listing)
	}
	return list, nil
}

func (db Database) AddListing(listing *models.Listing) error {
	var id int
	var createdAt string
	query := `INSERT INTO listings (mls_number, address, address_2, city, postal_code, state, country) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, created_at`
	err := db.Conn.QueryRow(query,
		listing.MLSNumber,
		listing.Address,
		listing.Address2,
		listing.City,
		listing.PostalCode,
		listing.State,
		listing.Country,
	).Scan(&id, &createdAt)
	if err != nil {
		return err
	}

	listing.ID = id
	listing.CreatedAt = createdAt
	return nil
}

func (db Database) GetListingById(listingId int) (models.Listing, error) {
	listing := models.Listing{}

	query := `SELECT * FROM listings WHERE id = $1;`
	row := db.Conn.QueryRow(query, listingId)
	switch err := row.Scan(&listing.ID, &listing.MLSNumber, &listing.Address, &listing.CreatedAt); err {
	case sql.ErrNoRows:
		return listing, ErrNoMatch
	default:
		return listing, err
	}
}

func (db Database) DeleteListing(listingId int) error {
	query := `DELETE FROM listings WHERE id = $1;`
	_, err := db.Conn.Exec(query, listingId)
	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}

func (db Database) UpdateListing(listingId int, listingData models.Listing) (models.Listing, error) {
	listing := models.Listing{}
	query := `UPDATE listings SET name=$1, description=$2 WHERE id=$3 RETURNING id, name, description, created_at;`
	err := db.Conn.QueryRow(query, listingData.MLSNumber, listingData.Address, listingId).Scan(&listing.ID, &listing.MLSNumber, &listing.Address, &listing.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return listing, ErrNoMatch
		}
		return listing, err
	}

	return listing, nil
}
