package main

import (
	"fetcher"
	"html/template"
	"net/http"
)

func indexHTML(w http.ResponseWriter, r *http.Request) {
	hnData := fetcher.GetContent(5)
	t, err := template.ParseFiles("index.html.j2")
	if err != nil {
		panic(err)
	}
	err = t.Execute(w, hnData)
	if err != nil {
		panic(err)
	}
}

func main() {

	http.HandleFunc("/", indexHTML)
	http.ListenAndServe(":8080", nil)

}
