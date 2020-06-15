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
}
