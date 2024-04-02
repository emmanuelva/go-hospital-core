package postgres

import (
	"context"
	"database/sql"
	"github.com/gofrs/uuid"
	"hospital-ma/internal/models"
	"hospital-ma/internal/repository"
	"log"
	"time"
)

const dbTimeout = time.Second * 3

type PGCountryStatesRepo struct {
	DB *sql.DB
}

func (u *PGCountryStatesRepo) AllCountryStates() (*models.CountryStates, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT id, name, alias, updated_at, created_at
				FROM country_states
				ORDER BY name`

	rows, err := u.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var countryStates models.CountryStates

	for rows.Next() {
		var countryState models.CountryState
		err := rows.Scan(
			&countryState.ID,
			&countryState.Name,
			&countryState.Alias,
			&countryState.UpdatedAt,
			&countryState.CreatedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		countryStates = append(countryStates, countryState)
	}
	return &countryStates, nil
}

func (u *PGCountryStatesRepo) InsertCountryState(countryState models.CountryState) (*models.CountryState, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newId, err = uuid.NewV4()
	var countryStateRow models.CountryState
	stmt := `insert into country_states (id, name, alias, updated_at, created_at)
        values ($1, $2, $3, $4, $5)
        returning id, name, alias, updated_at, created_at`
	err = u.DB.QueryRowContext(ctx, stmt,
		newId,
		countryState.Name,
		countryState.Alias,
		time.Now(),
		time.Now(),
	).Scan(
		&countryStateRow.ID,
		&countryStateRow.Name,
		&countryStateRow.Alias,
		&countryStateRow.UpdatedAt,
		&countryStateRow.CreatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &countryStateRow, nil
}

func NewCountryStatesRepo(DB *sql.DB) repository.CountryStatesRepo {
	return &PGCountryStatesRepo{DB: DB}
}
