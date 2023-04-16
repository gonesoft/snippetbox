package database

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestSnippetModel_Insert(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := regexp.QuoteMeta("INSERT INTO snippets (title, content, created_at, expires_at) VALUES($1, $2, NOW(), NOW() + INTERVAL '365 days') RETURNING id;")

	m := SnippetModel{
		DB: db,
	}

	mock.ExpectQuery(query).WithArgs("title", "content").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	id, err := m.Insert("title", "content", "")
	assert.NoError(t, err)
	assert.Equal(t, 1, id)

}
