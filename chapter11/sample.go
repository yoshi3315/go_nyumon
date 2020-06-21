package main

import "fmt"

func printArrayInt(a [3]int) {
	fmt.Println(a[0], a[1], a[2])
}

func printSliceInt(s []int) {
	if s == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("len: ", len(s), s[0], s[1])
	}
}

func main() {
	var array1 [3]int

	array2 := [3]int{20, 21}
	array3 := [...]int{30, 31, 32}
	array4 := [3]int{1: 41}

	printArrayInt(array1)
	printArrayInt(array2)
	printArrayInt(array3)
	printArrayInt(array4)

	var slice1 []int

	slice2 := []int{10, 20, 30}
	slice3 := []int{10, 20}

	printSliceInt(slice1)
	printSliceInt(slice2)
	printSliceInt(slice3)

	slice11 := make([]int, 2, 3)
	slice12 := append(slice11, 1)
	slice11[0] = 2

	slice13 := make([]int, 2, 2)
	slice14 := append(slice13, 1)
	slice13[0] = 2

	fmt.Println("len: ", len(slice11), len(slice12), len(slice13), len(slice14))
	fmt.Println("cap: ", cap(slice11), cap(slice12), cap(slice13), cap(slice14))
	fmt.Println("slice11: ", slice11)
	fmt.Println("slice12: ", slice12)
	fmt.Println("slice13: ", slice13)
	fmt.Println("slice14: ", slice14)

	slice15 := append(slice11, slice12...)
	fmt.Println("slice15: ", slice15)

	slice16 := make([]int, 6)
	copy(slice16, slice15)
	fmt.Println("slice16: ", slice16)

	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(s1[1:4], s1[1:], s1[:4], s1[:])

	a1 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(a1[1:4], a1[1:], a1[:4], a1[:])

	p1 := &a1
	fmt.Println(p1[1:4], p1[1:], p1[:4], p1[:])

	s2 := s1[1:4]
	s1[1] = 10
	s2[2] = 20
	fmt.Println(s1, s2)

	s3 := a1[1:4]
	a1[1] = 10
	s3[2] = 20
	fmt.Println(a1, s3)

	sliceA1 := []int{1, 2, 3, 4, 5}
	sliceA2 := sliceA1[0:2:2]
	sliceA3 := append(sliceA2, 6)
	fmt.Println("sliceA1: ", sliceA1)
	fmt.Println("sliceA2: ", sliceA2)
	fmt.Println("sliceA3: ", sliceA3)

	sliceB1 := []int{1, 2, 3, 4, 5}
	sliceB2 := sliceB1[0:2]
	sliceB3 := append(sliceB2, 6)
	fmt.Println("sliceB1: ", sliceB1)
	fmt.Println("sliceB2: ", sliceB2)
	fmt.Println("sliceB3: ", sliceB3)
}
