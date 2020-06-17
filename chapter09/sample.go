package main

import "fmt"

type myInt32 int32

type (
	myInt    int
	myString string
)

func main() {
	var n1 myInt32 = 1
	var n2 myInt32 = 2
	fmt.Println(n1 + n2)
}
