package main

import (
	"fmt"
)

func funcPanic1() {
	fmt.Println("funcPanic1 start")
	panic("funcpanic1 panic")
	fmt.Println("funcPanic1 end")
}

func funcPanic2() {
	fmt.Println("funcPanic2 start")
	funcPanic1()
	fmt.Println("funcPanic2 end")
}

func funcPanicDefer1() {
	defer fmt.Println("defer funcPanicDefer1")
	fmt.Println("funcPanicDefer1 panic start")
	panic("funcPanicDefer1 panic")
}

func funcPanicDefer2() {
	defer fmt.Println("defer funcPanicDefer2")
	funcPanicDefer1()
}

func funcPanic() {
	fmt.Println("funcPanic start")
}

func funcRecover() {
	fmt.Println("funcRecover start")
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("recover", r)
		}
	}()

	funcPanic()
	fmt.Println("funcRecover end")
}

func funcRuntimePanic() {
	var a []int
	fmt.Println(a[0])
}

func main() {
	// fmt.Println("main start")
	// funcPanic2()
	// fmt.Println("main end")

	// defer fmt.Println("defer main")
	// funcPanicDefer2()

	// fmt.Println("main start")
	// funcRecover()
	// fmt.Println("main end")

	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("recover", r)
		}
	}()
	funcRuntimePanic()
}
