package main

import "fmt"

func funcResultFunc() func() {
	n := 10

	return func() {
		fmt.Println(n)
	}
}

func main() {
	f := func(i, j int) int {
		return i + j
	}

	n := f(1, 2)
	fmt.Println(n)

	nn := 10
	ff := func() {
		fmt.Println(nn)
		nn++
	}

	ff()
	fmt.Println(nn)

	nn = 20
	ff()
	fmt.Println(nn)

	fff := funcResultFunc()
	fff()
}
