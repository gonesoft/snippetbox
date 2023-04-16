package main

import (
	"errors"
	"fmt"
	"github.com/gonesoft/snippetbox/pkg/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func (h *application) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	s, err := h.snippets.Latest()
	if err != nil {
		h.serverError(w, err)
		return
	}

	for _, snippet := range s {
		fmt.Fprintf(w, "%v", snippet)
	}

}

func (h *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
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

	files := []string{
		"ui/html/home.page.tmpl",
		"ui/html/base.layout.tmpl",
		"ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		h.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		h.serverError(w, err)
	}

	fmt.Fprintf(w, "%v", s)
}

func (h *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		h.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	title := "1 snail"
	content := "1 snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := "8"

	id, err := h.snippets.Insert(title, content, expires)
	if err != nil {
		h.serverError(w, err)
		return
	}
	if id == 0 {
		id = 1
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}

func (h *application) users(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific user..."))
}
