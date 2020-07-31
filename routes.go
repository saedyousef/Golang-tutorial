package main

import (
    "fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        title := vars["title"]
        page := vars["page"]

        fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
    })
	r.HandleFunc("/books/all/{title}", AllBooks).Methods("GET")
    http.ListenAndServe(":80", r)
}

func AllBooks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, vars["title"])
}