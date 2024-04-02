package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type City struct {
	ID           uuid.UUID    `json:"id"`
	CountryState CountryState `json:"country_state"`
	Name         string       `json:"name"`
	Alias        string       `json:"alias"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}

type Cities []City

func (u Cities) NameList() []string {
	var list []string
	for _, city := range u {
		list = append(list, city.Name)
	}
	return list
}
