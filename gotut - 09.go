package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, `<h1>Multiple lines example</h1>
			<p>Go is fast!</p>
			<p>...and simple!</p>
		`)

	fmt.Fprintf(w, "<h1>Hey There</h1>")
	fmt.Fprintf(w, "<p>You %s even add %s</p>", "can", "<strong>variables</strong>")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8000", nil)
}
