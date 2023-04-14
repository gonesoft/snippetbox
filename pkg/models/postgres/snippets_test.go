package postgres

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSnippetModel_Insert(t *testing.T) {
	db, err := sql.Open("postgres", "postgresql://postgres:password@localhost/snippetbox?sslmode=disable")
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	title := "2 snail"
	content := "2 snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"

	var id int64
	err = db.QueryRow("INSERT INTO snippets (title, content, created_at, expires_at) VALUES($1, $2, NOW(), NOW() + INTERVAL '365 days');",
		title,
		content).Scan(&id)
	assert.NoError(t, err)
	assert.Equal(t, 1, id)

}
