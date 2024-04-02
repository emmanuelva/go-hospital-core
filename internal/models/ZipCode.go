package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type ZipCode struct {
	ID        uuid.UUID `json:"id"`
	City      City      `json:"city"`
	Code      int       `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ZipCodes []ZipCode

func (u ZipCodes) CodeList() []int {
	var list []int
	for _, city := range u {
		list = append(list, city.Code)
	}
	return list
}
