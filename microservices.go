package main

import (
	"fmt"
	"net/http"
	"os"

	"api"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", echo)

	http.HandleFunc("/api/books", api.BooksHandleFunc) // to retrive books, store books.
	http.HandleFunc("/api/books/", api.BookHandleFunc) // to retrieve individual books by isbn, update them, delete them.

	http.ListenAndServe(port(), nil)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hola Cloud Go.")
}

func echo(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query()["message"][0]

	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, message)
}
