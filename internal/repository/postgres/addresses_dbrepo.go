package postgres

import (
	"context"
	"database/sql"
	"github.com/emmanuelva/go-hospital-core/internal/models"
	"github.com/emmanuelva/go-hospital-core/internal/repository"
	"github.com/emmanuelva/go-hospital-core/internal/utils"
	"github.com/gofrs/uuid"
	"log"
	"time"
)

type PGAddressesRepo struct {
	DB *sql.DB
}

func (u *PGAddressesRepo) AddressById(id string) (*models.Address, error) {
	return nil, nil
}

func (u *PGAddressesRepo) InsertAddress(address models.NewAddress) (*models.Address, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newId, err = uuid.NewV4()

	if err != nil {
		log.Fatal("Error generating UUID")
	}
	var addressRow models.Address
	stmt := `insert into addresses (id, zip_id, street, normalized_street, number, external_number, colony, normalized_colony, created_at, updated_at)
         values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
         returning id, street, normalized_street, number, external_number, colony, normalized_colony, created_at, updated_at`

	normalizedStreet, err := utils.NormalizeText(address.Street)
	if err != nil {
		return nil, err
	}

	normalizedColony, err := utils.NormalizeText(address.Colony)
	if err != nil {
		return nil, err
	}

	err = u.DB.QueryRowContext(ctx, stmt,
		newId,
		address.ZipCode.ID,
		address.Street,
		normalizedStreet,
		address.Number,
		address.ExternalNumber,
		address.Colony,
		normalizedColony,
		time.Now(),
		time.Now(),
	).Scan(
		&addressRow.ID,
		&addressRow.Street,
		&addressRow.NormalizedStreet,
		&addressRow.Number,
		&addressRow.ExternalNumber,
		&addressRow.Colony,
		&addressRow.NormalizedColony,
		&addressRow.CreatedAt,
		&addressRow.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &addressRow, nil
}

func NewAddressesRepo(DB *sql.DB) repository.AddressesRepo {
	return &PGAddressesRepo{DB: DB}
}
