package main

import (
	"fmt"
	"time"
)

func funcGo() {
	fmt.Println("go 1")

	time.Sleep(2 * time.Millisecond)

	fmt.Println("go 2")
}

func main() {
	go funcGo()

	time.Sleep(time.Millisecond)
	fmt.Println("main 1")
}
