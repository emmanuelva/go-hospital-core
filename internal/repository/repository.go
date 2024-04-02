package repository

import (
	"database/sql"
	"github.com/emmanuelva/go-hospital-core/internal/models"
	"github.com/gofrs/uuid"
)

type CountryStatesRepo interface {
	AllCountryStates() (*models.CountryStates, error)
	InsertCountryState(countryState models.CountryState) (*models.CountryState, error)
}

type CitiesRepo interface {
	CitiesByCountryState(countryState models.CountryState) (*models.Cities, error)
	InsertCity(city models.City) (*models.City, error)
}

type ZipCodesRepo interface {
	ZipCodesByCity(city models.City) (*models.ZipCodes, error)
	ZipCodeByNumber(zipCode int) (*models.ZipCode, error)
	InsertZipCode(zipCode models.ZipCode) (*models.ZipCode, error)
}

type AddressesRepo interface {
	AddressById(id string) (*models.Address, error)
	InsertAddress(address models.NewAddress) (*models.Address, error)
}

type PhonesRepo interface {
	InsertPhone(phone models.Phone) (*models.Phone, error)
}

type UsersRepo interface {
	AllUsers() ([]*models.User, error)
	InsertUser(user models.User) (*uuid.UUID, error)
}

type Repositories struct {
	Connection        *sql.DB
	CountryStatesRepo *CountryStatesRepo
	CitiesRepo        *CitiesRepo
	UsersRepo         *UsersRepo
	ZipCodesRepo      *ZipCodesRepo
	AddressesRepo     *AddressesRepo
	PhonesRepo        *PhonesRepo
}
