package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type Snippet struct {
	ID         int
	Title      string
	Content    string
	Created_at time.Time
	Expires_at time.Time
}

type User struct {
	ID       int
	Name     string
	Email    string
	Password []byte
	Created  time.Time
	Active   bool
}
