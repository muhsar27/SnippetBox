package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from SnippetBox"))
}

// Displays a snippet
func snippetView(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)
	w.Write([]byte(msg))
}

// Displays a form for creating a snippet
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating something something"))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	//201 represents created
	w.Header().Add("Server", "Go")

	w.Write([]byte("Saves a new snippet"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", home)
	//including url parameters
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)
	//specifying the http method
	log.Println("Server starting on :4000")

	//ListenAndServe starts our server and binds it to the address specified
	err := http.ListenAndServe(":4000", mux)
	//logs an error if err is non-nil
	log.Fatal(err)
}
