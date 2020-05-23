package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type NewsAggPage struct {
	Title string
	News  map[string]string
	Arr   [4]int
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Hello world </h1")
}

func newAggHandler(w http.ResponseWriter, r *http.Request) {

	newsMap := make(map[string]string)
	newsMap["title1"] = "keyword1"
	newsMap["title2"] = "keyword2"

	intArray := [...]int{3, 4, 5, 6}

	fmt.Println(newsMap)

	p := NewsAggPage{Title: "Amazing News Aggregator", News: newsMap, Arr: intArray}
	t, _ := template.ParseFiles("basictemplating.html")
	err := t.Execute(w, p)
	fmt.Println(err)
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/agg/", newAggHandler)
	http.ListenAndServe(":8000", nil)
}
