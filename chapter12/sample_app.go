package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type serverHTML struct {
	Title string
	Body  string
}

var tmpl = template.Must(template.ParseFiles("html.tmpl"))

func (s *serverHTML) serverHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if err := tmpl.Execute(w, s); err != nil {
		panic(err)
	}
}

func main() {
	helloHTML := &serverHTML{"Hello Title", "Hello"}
	goodbyeHTML := &serverHTML{"Goodbye Title", "Goodbye"}
	http.Handle("/hello", helloHTML)
	http.Handle("/goodbye", goodbyeHTML)

	if err := http.ListenAndServe(":4000", nil); err != nil {
		fmt.Println(err)
	}
}
