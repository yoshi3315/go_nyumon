package main

import (
	"html/template"
	"net/http"
	"strconv"
)

const maru, batsu = "○", "✕"

type Board [3][3]string

type ViewData struct {
	Turn   string
	Board  *Board
	Win    bool
	Draw   bool
	Winner string
}

var tmpl *template.Template = template.Must(template.ParseFiles("game.tmpl"))

func (v *ViewData) Execute(w http.ResponseWriter) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	if err := tmpl.Execute(w, v); err != nil {
		panic(err)
	}
}

func gameHandle(w http.ResponseWriter, r *http.Request) {
	turn, nextTurn := turnFormValue(r)
	board := boardFormValue(r)

	win, draw, winner := false, false, ""

	if turn != "" {
		win = board.win(turn)

		if win {
			winner = turn
			board.setBar()
		} else {
			draw = board.draw()
		}
	}

	v := ViewData{nextTurn, board, win, draw, winner}
	v.Execute(w)
}

var winBoardIndexArray = [...][3]struct{ row, col int }{
	{{0, 0}, {0, 1}, {0, 2}},
	{{1, 0}, {1, 1}, {1, 2}},
	{{2, 0}, {2, 1}, {2, 2}},

	{{0, 0}, {1, 0}, {2, 0}},
	{{0, 1}, {1, 1}, {2, 1}},
	{{0, 2}, {1, 2}, {2, 2}},

	{{0, 0}, {1, 1}, {2, 2}},
	{{0, 2}, {1, 1}, {2, 0}},
}

func (b *Board) win(turn string) bool {
	for _, w := range winBoardIndexArray {
		if (b[w[0].row][w[0].col] == turn) && (b[w[1].row][w[1].col] == turn) && (b[w[2].row][w[2].col] == turn) {
			return true
		}
	}
	return false
}

func (b *Board) draw() bool {
	for row, rows := range b {
		for col, _ := range rows {
			if b[row][col] == "" {
				return false
			}
		}
	}
	return true
}

func (b *Board) setBar() {
	for row, rows := range b {
		for col, _ := range rows {
			if b[row][col] == "" {
				b[row][col] = "-"
			}
		}
	}
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
