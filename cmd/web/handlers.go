package main

import (
	"errors"
	"fmt"
	"github.com/gonesoft/snippetbox/pkg/forms"
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
	h.render(w, r, "create.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (h *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		h.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("title", "content", "expires")
	form.MaxLength("title", 100)
	form.PermittedValues("expires", "365", "7", "1")

	if !form.Valid() {
		h.render(w, r, "create.page.tmpl", &templateData{Form: form})
		return
	}

	id, err := h.snippets.Insert(form.Get("title"), form.Get("content"), form.Get("expires"))
	if err != nil {
		h.serverError(w, err)
		return
	}

	h.session.Put(r, "flash", "Snippet successfully created!")

	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}

func (h *application) users(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific user..."))
}
