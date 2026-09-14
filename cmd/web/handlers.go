package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
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

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// Displays a form for creating a snippet
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating something something"))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	//201 represents created
	w.WriteHeader(http.StatusCreated)

	fmt.Fprint(w, "Saves a new snippet")
}
