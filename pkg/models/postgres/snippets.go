package postgres

import (
	"database/sql"
	"github.com/gonesoft/snippetbox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

func NewSnippetModel(db *sql.DB) *SnippetModel {
	return &SnippetModel{DB: db}
}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	var id int
	err := m.DB.QueryRow("INSERT INTO snippets (title, content, created_at, expires_at) VALUES($1, $2, NOW(), NOW() + INTERVAL '365 days');",
		title,
		content).Scan(&id)
	if err != nil {
		return 0, nil
	}

	return id, nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
