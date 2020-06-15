package main

import "fmt"

func main() {
	if x := 3; x == 1 {
		fmt.Println("x == 1")
	} else if x == 2 {
		fmt.Println("x == 2")
	} else {
		fmt.Println(x)
	}

	x := 1
	switch y := x + 1; y {
	case 1:
		z := x + y
		fmt.Println("case 1:", x, y, z)
	case 2:
		fmt.Println("case 2:", x, y)
	}

	switch x {
	case 1:
		fmt.Println("A")
		fallthrough
	case 2:
		fmt.Println("B")
	case 3:
		fmt.Println("C")
	}

	y := 2
	switch x {
	case 1:
		fmt.Println("A")
		if y == 2 {
			break
		}

		fmt.Println("B")
	}

	n := 0
	for i := 1; i <= 10; i++ {
		n += i
	}
	fmt.Println(n)

	for i := 1; i <= 10; i++ {
		if (i % 2) == 0 {
			continue
		}

		fmt.Println(i, ",")
	}

	fmt.Println()

	for i := 1; i <= 10; i++ {
		if i == 5 {
			break
		}
		fmt.Print(i, ",")
	}

OuterLoop:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println(i, j)

			if (i == 0) && (j == 1) {
				fmt.Println("continue")
				continue OuterLoop
			}

			if (i == 1) && (j == 1) {
				fmt.Println("break")
				break OuterLoop
			}
		}
	}

	m := 1
	for m <= 1000 {
		m *= 2
	}

	fmt.Println(m)

	o := 1
	for {
		o *= 2
		if 1000 <= o {
			break
		}
	}

	fmt.Println(o)

	s := "aあいう"

	for i, r := range s {
		fmt.Println(i, r, string(r))
	}

	goto label

	fmt.Println("実行されない")

label:
	fmt.Println("goto文によるジャンプ")
}
