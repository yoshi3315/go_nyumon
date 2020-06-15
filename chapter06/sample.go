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
}
