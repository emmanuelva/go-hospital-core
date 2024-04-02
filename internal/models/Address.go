package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type Address struct {
	ID               uuid.UUID `json:"id"`
	ZipCode          ZipCode   `json:"zip_code"`
	Street           string    `json:"street"`
	NormalizedStreet string    `json:"-"`
	Number           string    `json:"number"`
	ExternalNumber   string    `json:"external_number"`
	Colony           string    `json:"colony"`
	NormalizedColony string    `json:"normalized_colony"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type NewAddress struct {
	ID             uuid.UUID `json:"id"`
	ZipCode        ZipCode   `json:"zip_code"`
	Street         string    `json:"street"`
	Number         string    `json:"number"`
	ExternalNumber string    `json:"external_number"`
	Colony         string    `json:"colony"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
