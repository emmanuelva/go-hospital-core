package postgres

import (
	"context"
	"database/sql"
	"github.com/emmanuelva/go-hospital-core/internal/models"
	"github.com/emmanuelva/go-hospital-core/internal/repository"
	"github.com/gofrs/uuid"
	"log"
	"time"
)

type PGCitiesRepo struct {
	DB *sql.DB
}

func (u *PGCitiesRepo) CitiesByCountryState(countryState models.CountryState) (*models.Cities, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT id, name, alias, updated_at, created_at
				FROM cities
				WHERE country_state_id = $1
				ORDER BY name`

	rows, err := u.DB.QueryContext(ctx, query, countryState.ID.String())

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cities models.Cities

	for rows.Next() {
		var city models.City
		err := rows.Scan(
			&city.ID,
			&city.Name,
			&city.Alias,
			&city.UpdatedAt,
			&city.CreatedAt,
		)

		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}
		city.CountryState = countryState

		cities = append(cities, city)
	}

	return &cities, nil
}

func (u *PGCitiesRepo) InsertCity(city models.City) (*models.City, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newId, err = uuid.NewV4()
	var cityRow models.City
	stmt := `insert into cities (id, country_state_id, name, alias, updated_at, created_at)
        values ($1, $2, $3, $4, $5, $6)
        returning id, name, alias, updated_at, created_at`
	err = u.DB.QueryRowContext(ctx, stmt,
		newId,
		city.CountryState.ID,
		city.Name,
		city.Alias,
		time.Now(),
		time.Now(),
	).Scan(
		&cityRow.ID,
		&cityRow.Name,
		&cityRow.Alias,
		&cityRow.UpdatedAt,
		&cityRow.CreatedAt,
	)

	if err != nil {
		return nil, err
	}
	cityRow.CountryState = city.CountryState
	return &cityRow, nil
}

func NewCitiesRepo(DB *sql.DB) repository.CitiesRepo {
	return &PGCitiesRepo{DB: DB}
}
