package mock

import (
	"github.com/gonesoft/snippetbox/pkg/models"
	"time"
)

var mockSnippet = &models.Snippet{
	ID:         1,
	Title:      "an old silent pond",
	Content:    "an old silent pond...",
	Created_at: time.Time{},
	Expires_at: time.Time{},
}

type SnippetModel struct{}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 2, nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	switch id {
	case 1:
		return mockSnippet, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return []*models.Snippet{mockSnippet}, nil
}
