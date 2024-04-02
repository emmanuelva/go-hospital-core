package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type CountryState struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Alias     string    `json:"alias"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CountryStates []CountryState

func (u CountryStates) NameList() []string {
	var list []string
	for _, countryState := range u {
		list = append(list, countryState.Name)
	}
	return list
}
