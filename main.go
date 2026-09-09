package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from SnippetBox"))
}

// Displays a snippet
func snippetview(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Viewing Something something"))
}

// Displays a form for creating a snippet
func snippetcreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating something something"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetview)
	mux.HandleFunc("/snippet/create", snippetcreate)

	log.Println("Server starting on :4000")

	//ListenAndServe starts our server and binds it to the address specified
	err := http.ListenAndServe(":4000", mux)
	//logs an error if err is non-nil
	log.Fatal(err)
}
