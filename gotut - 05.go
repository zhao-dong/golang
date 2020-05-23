package main

import (
	"fmt"
	"net/http"
)

func index_handler(W http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(W, "Whoa, Go is neat!")
}

func about_handler(W http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(W, "Expert web design by Dong!")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)
	http.ListenAndServe(":8000", nil)
}
