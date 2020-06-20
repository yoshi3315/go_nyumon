package main

import (
	"fmt"
	"unsafe"
)

type myStruct struct {
	field int
}

type point struct {
	x, y int
}

func main() {
	var v int32 = 100
	var pointer *int32 = &v
	fmt.Println("ポインタの値: ", pointer)

	fmt.Println("関節演算子の結果: ", *pointer)

	v = 200
	fmt.Println("変数vの更新: ", *pointer)

	*pointer = 300
	fmt.Println("ポインタから値を更新: ", v)

	fmt.Println("&myStruct{1}: ", &myStruct{1})

	s := myStruct{1}
	fmt.Println("&s.field: ", &s.field)

	array := [10]int{}
	fmt.Println("&array[0]: ", &array[0])

	slice := []int{0}
	fmt.Println("&slice[0]: ", &slice[0])

	n := new(int)
	fmt.Println(n, *n)

	p1 := new(point)
	fmt.Println(*p1)

	p2 := &point{}
	fmt.Println(*p2)

	s1 := myStruct{1}
	s2 := s1
	s2.field = 2
	fmt.Println(s1, s2)

	s3 := &s1

	s3.field = 3
	fmt.Println(s1, *s3)

	pointerInt1 := new(int)
	unsafePointer1 := unsafe.Pointer(pointerInt1)
	fmt.Println(unsafePointer1)

	u := uintptr(unsafePointer1)
	u += 2

	unsafePointer1 = unsafe.Pointer(u)
	fmt.Println(unsafePointer1)

	pointerInt2 := new(int)
	unsafePointer2 := unsafe.Pointer(pointerInt2)
	pointerString := (*string)(unsafePointer2)

	fmt.Println(*pointerString)
}
