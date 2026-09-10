package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

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
