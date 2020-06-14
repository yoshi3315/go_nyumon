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

var (
	n1 = 6789
	n2 = 04567
	n3 = 0xCDEF

	f1 = 1.2
	f2 = 1.2e+3

	i2 = 1.2i
	i3 = 3 + 1.2i

	r1 = 'a'
	r2 = 'あ'

	b1 byte = 1.0
	b2 byte = 0i

	s1 = "abc¥nあいうえお¥nかきくけこ"
	s2 = `abc
	あいうえお
	かきくけこ`

	s3 = "¥u3042"
	s4 = "¥xE3¥x81¥x82"
)

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

	fmt.Println(n1, n2, n3, f1, f2, i2, i3)
	fmt.Println(r1, r2, b1, b2)
	fmt.Println(s1, s2, s3, s4)
}

func f() {
	n = 2
}
