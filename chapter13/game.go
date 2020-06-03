package main

import (
	"html/template"
	"net/http"
)

const maru, batsu = "○", "✕"

type Board [3][3]string

type ViewData struct {
	Turn  string
	Board *Board
}

var tmpl *template.Template = template.Must(template.ParseFiles("game.tmpl"))

func (v *ViewData) Execute(w http.ResponseWriter) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	if err := tmpl.Execute(w, v); err != nil {
		panic(err)
	}
}

func gameHandle(w http.ResponseWriter, r *http.Request) {
	turn := maru
	board := &Board{}
	v := ViewData{turn, board}
	v.Execute(w)
}

func main() {
	http.HandleFunc("/game", gameHandle)

	if err := http.ListenAndServe(":4000", nil); err != nil {
		panic(err)
	}
}
