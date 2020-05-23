package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type NewsAggPage struct {
	Title string
	News  string
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Hello world </h1")
}

func newAggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "Amazing News Aggregator", News: "some news"}
	t, _ := template.ParseFiles("basictemplating.html")
	err := t.Execute(w, p)
	fmt.Println(err)
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/agg/", newAggHandler)
	http.ListenAndServe(":8000", nil)
}
