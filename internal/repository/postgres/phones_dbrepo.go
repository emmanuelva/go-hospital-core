package postgres

import (
	"context"
	"database/sql"
	"github.com/emmanuelva/go-hospital-core/internal/models"
	"github.com/emmanuelva/go-hospital-core/internal/repository"
	"github.com/gofrs/uuid"
	"log"
	"regexp"
	"time"
)

type PGPhonesRepo struct {
	DB *sql.DB
}

func (u *PGPhonesRepo) InsertPhone(phone models.Phone) (*models.Phone, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newId, err = uuid.NewV4()

	if err != nil {
		log.Fatal("Error generating UUID")
	}

	var phoneRow models.Phone

	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	normalizedPhoneBytes := re.FindAll([]byte(phone.Number), -1)
	normalizedPhone := string(normalizedPhoneBytes[0])

	stmt := `insert into phones (id, number, extension, phone_type, created_at, updated_at)
			values ($1, $2, $3, $4, $5, $6)
			returning id, number, extension, phone_type, created_at, updated_at`

	err = u.DB.QueryRowContext(ctx, stmt,
		newId,
		normalizedPhone,
		phone.Extension,
		phone.PhoneType,
		time.Now(),
		time.Now(),
	).Scan(
		&phoneRow.ID,
		&phoneRow.Number,
		&phoneRow.Extension,
		&phoneRow.PhoneType,
		&phoneRow.CreatedAt,
		&phoneRow.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &phoneRow, nil
}

func NewPhonesRepo(DB *sql.DB) repository.PhonesRepo {
	return &PGPhonesRepo{DB: DB}
}
