package repositories

import (
	"database/sql"
	"github.com/emmanuelva/go-hospital-core/internal/repository"
	postgres2 "github.com/emmanuelva/go-hospital-core/internal/repository/postgres"
	"time"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3

func InitPostgresRepo(DB *sql.DB) repository.Repositories {
	countryStateRepo := postgres2.NewCountryStatesRepo(DB)
	citiesRepo := postgres2.NewCitiesRepo(DB)
	usersRepo := postgres2.NewUsersRepo(DB)
	zipCodesRepo := postgres2.NewZipCodesRepo(DB)
	addressesRepo := postgres2.NewAddressesRepo(DB)
	phonesRepo := postgres2.NewPhonesRepo(DB)
	return repository.Repositories{
		Connection:        DB,
		CountryStatesRepo: &countryStateRepo,
		CitiesRepo:        &citiesRepo,
		UsersRepo:         &usersRepo,
		ZipCodesRepo:      &zipCodesRepo,
		AddressesRepo:     &addressesRepo,
		PhonesRepo:        &phonesRepo,
	}
}

func (u *PostgresDBRepo) Connection() *sql.DB {
	return u.DB
}
