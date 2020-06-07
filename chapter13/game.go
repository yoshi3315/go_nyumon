package main

import (
	"html/template"
	"net/http"
	"strconv"
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
	_, nextTurn := turnFormValue(r)
	board := boardFormValue(r)

	v := ViewData{nextTurn, board}
	v.Execute(w)
}

func boardFormValue(r *http.Request) *Board {
	var board Board
	for row, rows := range board {
		for col, _ := range rows {
			name := "c" + strconv.Itoa(row) + strconv.Itoa(col)
			board[row][col] = r.FormValue(name)
		}
	}

	return &board
}

var nextTurnMap = map[string]string{
	maru:  batsu,
	batsu: maru,
	"":    maru,
}

func turnFormValue(r *http.Request) (string, string) {
	turn := r.FormValue("turn")
	nextTurn := nextTurnMap[turn]
	return turn, nextTurn
}

func main() {
	http.HandleFunc("/game", gameHandle)

	if err := http.ListenAndServe(":4000", nil); err != nil {
		panic(err)
	}
}
