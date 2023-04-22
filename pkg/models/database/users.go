package database

import (
	"database/sql"
	"errors"
	"github.com/gonesoft/snippetbox/pkg/models"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email, password string) error {
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	_, err = m.DB.Exec(`INSERT INTO users (name, email, hashed_password, created) VALUES ($1, $2, $3, $4)`,
		name, email, hasedPassword, time.Now())

	if err != nil {
		if pgError, ok := err.(*pq.Error); ok && pgError.Code == "23505" {
			return models.ErrDuplicateEmail
		}
		return err
	}

	return nil
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	var (
		id             int
		hashedPassword []byte
	)

	err := m.DB.QueryRow(`SELECT id, hashed_password FROM users WHERE email = $1 and active = TRUE;`, email).Scan(
		&id, &hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, models.ErrInvalidCredentials
		}
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return 0, models.ErrInvalidCredentials
		}
		return 0, err
	}

	return id, nil
}

func (m *UserModel) Get(id int) (*models.User, error) {
	usr := &models.User{}

	err := m.DB.QueryRow(`SELECT id, name, email, created, active FROM users WHERE id = $1`, id).Scan(
		&usr.ID, &usr.Name, &usr.Email, &usr.Created, &usr.Active)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return usr, nil
}
