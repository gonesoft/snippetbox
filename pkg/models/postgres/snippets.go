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
	s := &models.Snippet{}
	err := m.DB.QueryRow("SELECT id, title, content, created_at, expires_at FROM snippets WHERE expires_at > NOW() AND id = $1", id).Scan(
		&s.ID,
		&s.Title,
		&s.Content,
		&s.Created_at,
		&s.Expires_at,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, models.ErrNoRecord
		}
		return nil, err
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	rows, err := m.DB.Query("SELECT id, title, content, created_at, expires_at FROM snippets WHERE expires_at > NOW() ORDER BY created_at DESC LIMIT 10")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, models.ErrNoRecord
		}
		return nil, err
	}
	defer rows.Close()

	snippets := []*models.Snippet{}
	for rows.Next() {
		s := &models.Snippet{}
		err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created_at, &s.Expires_at)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
