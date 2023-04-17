package main

import (
	"errors"
	"fmt"
	"github.com/gonesoft/snippetbox/pkg/models"
	"net/http"
	"strconv"
)

func (h *application) showHome(w http.ResponseWriter, r *http.Request) {
	s, err := h.snippets.Latest()
	if err != nil {
		h.serverError(w, err)
		return
	}

	data := &templateData{Snippets: s}

	h.render(w, r, "home.page.tmpl", data)

}

func (h *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		h.notFound(w)
		return
	}

	s, err := h.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			h.notFound(w)
		} else {
			h.serverError(w, err)
		}
		return
	}

	data := &templateData{Snippet: s}

	h.render(w, r, "show.page.tmpl", data)
}

func (h *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display the create snippet form..."))
}

func (h *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	title := "1 snail"
	content := "1 snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := "8"

	id, err := h.snippets.Insert(title, content, expires)
	if err != nil {
		h.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}

func (h *application) users(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific user..."))
}
