package main

import (
	"fmt"
	"sync"
)

var n1 = 0

var sharedVar string

var mu sync.Mutex

func writeSharedVar(s string) {
	mu.Lock()
	defer mu.Unlock()
	sharedVar = s
}

func readSharedVar() string {
	mu.Lock()
	defer mu.Unlock()
	return sharedVar
}

func main() {
	n2 := 0

	ch := make(chan struct{})

	go func() {
		n1, n2 = 1, 2
		fmt.Println("go: ", n1, &n1, n2, &n2)
		close(ch)
	}()

	<-ch
	fmt.Println("main: ", n1, &n1, n2, &n2)

	s1 := "AB"
	s2 := "C"
	fmt.Println("s1: ", s1, len(s1))
	fmt.Println("s2: ", s2, len(s2))

	go func() {
		for {
			writeSharedVar(s1)
			writeSharedVar(s2)
		}
	}()

	for {
		s3 := readSharedVar()

		if len(s1) == len(s3) && s2[0] == s3[0] {
			fmt.Println("s3: ", s3, len(s3))
			break
		}
	}
}
