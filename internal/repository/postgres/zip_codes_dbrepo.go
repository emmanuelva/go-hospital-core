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

type PGZipCodesRepo struct {
	DB *sql.DB
}

func (u *PGZipCodesRepo) ZipCodesByCity(city models.City) (*models.ZipCodes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT id, code, updated_at, created_at
				FROM zip_codes
				WHERE city_id = $1
				ORDER BY code`

	rows, err := u.DB.QueryContext(ctx, query, city.ID.String())

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var zipCodes models.ZipCodes

	for rows.Next() {
		var zipCode models.ZipCode
		err := rows.Scan(
			&zipCode.ID,
			&zipCode.Code,
			&zipCode.UpdatedAt,
			&zipCode.CreatedAt,
		)

		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}
		zipCode.City = city
		zipCodes = append(zipCodes, zipCode)
	}
	return &zipCodes, nil
}

func (u *PGZipCodesRepo) ZipCodeByNumber(zipCode int) (*models.ZipCode, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT id, code, updated_at, created_at
				FROM zip_codes
				WHERE code = $1
				ORDER BY code`

	var zipCodeResult models.ZipCode

	row := u.DB.QueryRowContext(ctx, query, zipCode)
	err := row.Scan(
		&zipCodeResult.ID,
		&zipCodeResult.Code,
		&zipCodeResult.UpdatedAt,
		&zipCodeResult.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &zipCodeResult, nil
}

func (u *PGZipCodesRepo) InsertZipCode(zipCode models.ZipCode) (*models.ZipCode, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newId, err = uuid.NewV4()
	var zipCodeRow models.ZipCode
	stmt := `insert into zip_codes (id, city_id, code, created_at, updated_at)
        values ($1, $2, $3, $4, $5)
        returning id, code, created_at, updated_at`
	err = u.DB.QueryRowContext(ctx, stmt,
		newId,
		zipCode.City.ID,
		zipCode.Code,
		time.Now(),
		time.Now(),
	).Scan(
		&zipCodeRow.ID,
		&zipCodeRow.Code,
		&zipCodeRow.UpdatedAt,
		&zipCodeRow.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	zipCodeRow.City = zipCode.City

	return &zipCodeRow, nil
}

func NewZipCodesRepo(DB *sql.DB) repository.ZipCodesRepo {
	return &PGZipCodesRepo{DB: DB}
}
