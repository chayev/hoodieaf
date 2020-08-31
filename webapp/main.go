package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Listening on :8080...")
	http.HandleFunc("/", handler)
	http.HandleFunc("/results", viewResults)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("tpl/base.html")
	t.Execute(w, nil)
}

func viewResults(w http.ResponseWriter, r *http.Request) {

	var input string

	for _, n := range r.URL.Query()["input"] {
		input = n
	}

	var output string

	if input == "" {
		output = "Enter an input."
	}
	// else {
	// 	output = lib.Create(input)
	// }

	t, _ := template.ParseFiles("tpl/results.html")
	t.Execute(w, &output)
}
