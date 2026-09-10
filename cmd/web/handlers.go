package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	files := []string{
		"ui/html/base.tmpl.html",
		"ui/html/partials/nav.tmpl.html",
		"ui/html/pages/home.tmpl.html",	
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		log.Println(err.Error())
		//Error parsing
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)

	if err != nil {
		log.Print(err.Error())
		//Error executing
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	w.Write([]byte("Hello from SnippetBox"))
}

// Displays a snippet
func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// Displays a form for creating a snippet
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating something something"))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	//201 represents created
	w.WriteHeader(http.StatusCreated)

	fmt.Fprint(w, "Saves a new snippet")
}
