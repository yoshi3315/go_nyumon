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

func funcSend(ch chan<- int) {
	fmt.Println("送信 10")
	ch <- 10

	for n := 0; n < 3; n++ {
		time.Sleep(time.Millisecond)
		fmt.Println("送信 ", n)
		ch <- n
	}
}

func funcRecv(ch <-chan int) {
	for n := 0; n < 3; n++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Println("受信 ", <-ch)
	}
}

func funcSendClose(ch chan<- int) {
	for n := 0; n < 3; n++ {
		ch <- n
	}

	close(ch)
}

func funcSendFor(ch chan<- int) {
	for n := 0; n < 3; n++ {
		ch <- n
	}

	close(ch)
}

func funcSelect(ch1 chan<- string, ch2 <-chan string) {
	select {
	case ch1 <- "a":
		fmt.Println("select ch1: a")
	case n := <-ch2:
		fmt.Println("select ch2: ", n)
	}
}

func funcSelectDefault(ch1 chan<- string, ch2 <-chan string) {
	select {
	case ch1 <- "a":
		fmt.Println("select ch1: a")
	case n := <-ch2:
		fmt.Println("select ch2: ", n)
	default:
		fmt.Println("select default")
	}
}

func main() {
	go funcGo()

	time.Sleep(time.Millisecond)
	fmt.Println("main 1")

	ch := make(chan int)
	go funcSend(ch)
	n := <-ch
	fmt.Println("受信 ", n)

	fmt.Println("バッファ3")
	ch1 := make(chan int, 3)
	go funcSend(ch1)
	funcRecv(ch1)

	fmt.Println("バッファ1")
	ch2 := make(chan int, 1)
	go funcSend(ch2)
	funcRecv(ch2)

	fmt.Println("バッファ0")
	ch3 := make(chan int)
	go funcSend(ch3)
	funcRecv(ch3)

	ch4 := make(chan int)
	go funcSendClose(ch4)

	for {
		n, ok := <-ch4

		if !ok {
			break
		}

		fmt.Println("受信 ", n)
	}

	fmt.Println("クローズ ", <-ch4)

	ch5 := make(chan int)
	go funcSendFor(ch5)

	for n := range ch5 {
		fmt.Println("受信 ", n)
	}

	ch6 := make(chan string)
	ch7 := make(chan string)

	go func() { <-ch6 }()
	funcSelect(ch6, ch7)

	go func() { ch7 <- "b" }()
	funcSelect(ch6, ch7)

	funcSelectDefault(ch6, ch7)
}
