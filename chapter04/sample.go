package main

import (
	"fmt"
	"strconv"
)

var (
	x int
	y int = 0
	z     = 0
)

var x1, x2 int
var y1, y2 int = 1, 2
var z1, z2 = 1, "a"

var i1, err1 = strconv.Atoi("10")
var _, err2 = strconv.Atoi("10")

var n = 1

func main() {
	x3 := 1
	y3, z3 := 2, 3

	fmt.Println(x, y, z)
	fmt.Println(x3, y3, z3)

	fmt.Println(n)
	f()
	fmt.Println(n)

	p := 1

	if p == 1 {
		q := 2
		fmt.Println(p, q)
	}
}

func f() {
	n = 2
}
