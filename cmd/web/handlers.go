package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"snippetbox.alexedwards.net/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	files := []string{
		"ui/html/base.tmpl.html",
		"ui/html/partials/nav.tmpl.html",
		"ui/html/pages/home.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)

	if err != nil {
		app.serverError(w, r, err)
	}
	w.Write([]byte("Hello from SnippetBox"))
}

// Displays a snippet
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	//http.ServeFile(w, r, "./ui/html/base.tmpl.html")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord){
			http.NotFound(w, r)
		} else{
			app.serverError(w , r , err)
		}
		return 
	}

	fmt.Fprintf(w, "%+v", snippet)
}

// Displays a form for creating a snippet
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating something something"))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	title := "0 snail"
	content := "0 snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := 7 

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return 
	}
	http.Redirect(w, r , fmt.Sprintf("/snippet/view/%d", id),http.StatusSeeOther)

	//201 represents created
	// w.WriteHeader(http.StatusCreated)

	// fmt.Fprint(w, "Saves a new snippet")
}
