package models

import (
	"fmt"
	"net/http"
)

type Listing struct {
	ID int `json:"id"`
	MLSNumber int `json:"mls_number"`
	Address string `json:"address"`
	Address2 *string `json:"address_2"`
	City string `json:"city"`
	PostalCode string `json:"postal_code"`
	State string `json:"state"`
	Country string `json:"country"`
	CreatedAt string `json:"created_at"`
}

type ListingList struct {
	Listings []Listing `json:"items"`
}

func (i *Listing) Bind(r *http.Request) error {
	if i.Address == "" {
		return fmt.Errorf("address is a required field")
	}
	if i.PostalCode == "" {
		return fmt.Errorf("postal_code is a required field")
	}
	if i.State == "" {
		return fmt.Errorf("state is a required field")
	}
	if i.Country == "" {
		return fmt.Errorf("country is a required field")
	}
	return nil
}

func (*ListingList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Listing) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
