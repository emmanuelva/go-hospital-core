package postgres

import (
	"context"
	"database/sql"
	"github.com/emmanuelva/go-hospital-core/internal/models"
	"github.com/emmanuelva/go-hospital-core/internal/repository"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type PGUsersRepo struct {
	DB *sql.DB
}

func (u *PGUsersRepo) AllUsers() ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT * 
				FROM users
				ORDER BY last_name`

	rows, err := u.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.Password,
			&user.Active,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		users = append(users, &user)
	}
	return users, nil
}

func (u *PGUsersRepo) InsertUser(user models.User) (*uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return nil, nil
	}
	var newID uuid.UUID
	stmt := `insert into users (email, first_name, last_name, password, active, created_at, updated_at)
        values ($1, $2, $3, $4, $5, $6, $7) returning id`
	err = u.DB.QueryRowContext(ctx, stmt,
		user.Email,
		user.FirstName,
		user.LastName,
		hashedPassword,
		user.Active,
		time.Now(),
		time.Now(),
	).Scan(&newID)
	if err != nil {
		return nil, err
	}
	return &newID, nil
}

func NewUsersRepo(DB *sql.DB) repository.UsersRepo {
	return &PGUsersRepo{DB: DB}
}
