package main

import (
	"fmt"
	"runtime"
	"time"
)

const c = 10000000

func funcGoLoop(ch chan<- int) {
	n := 0
	for i := 0; i < c; i++ {
		n += i
	}
	ch <- n
}

func funcGo() {
	ch := make(chan int)

	go funcGoLoop(ch)
	go funcGoLoop(ch)

	<-ch
	<-ch
}

func main() {
	runtime.GOMAXPROCS(1)
	t1 := time.Now()
	funcGo()
	fmt.Println("GOMAXPROCS=1: ", time.Since(t1))

	runtime.GOMAXPROCS(2)
	t2 := time.Now()
	funcGo()
	fmt.Println("GOMAXPROCS=2: ", time.Since(t2))

	runtime.GOMAXPROCS(1)
	go func() {
		fmt.Println("go")

		for {

		}
	}()

	time.Sleep(time.Millisecond)
	fmt.Println("main")
}
