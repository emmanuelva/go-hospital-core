package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type PhoneType string

const (
	PhoneHome   PhoneType = "HOME"
	PhoneOffice           = "OFFICE"
	PhoneCell             = "CELL"
)

type Phone struct {
	ID        uuid.UUID `json:"id"`
	Number    string    `json:"number"`
	Extension string    `json:"extension"`
	PhoneType PhoneType `json:"phone_type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
