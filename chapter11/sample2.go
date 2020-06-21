package main

import "fmt"

type myStruct struct {
	s string
	x int
}

func main() {
	s1 := [...]myStruct{myStruct{"A", 1}, myStruct{"B", 1}}
	s2 := [...]myStruct{{"A", 1}, {"B", 1}}

	array1 := [...][2]int{[2]int{1, 2}, [2]int{3, 4}}
	array2 := [...][2]int{{1, 2}, {3, 4}}

	slice1 := [...][]int{[]int{1, 2}, []int{3, 4}}
	slice2 := [...][]int{{1, 2}, {3, 4}}

	map1 := [...]map[int]int{map[int]int{1: 2}}
	map2 := [...]map[int]int{{1: 2}}

	fmt.Println(s1, s2)
	fmt.Println(array1, array2)
	fmt.Println(slice1, slice2)
	fmt.Println(map1, map2)

	array := []string{"a", "b", "c"}
	for i, e := range array {
		fmt.Print(i, ":", e, ",")
	}
	fmt.Println()

	slice := []string{"d", "e", "f"}
	for i, e := range slice {
		fmt.Print(i, ":", e, ",")
	}
	fmt.Println()

	for i := range array {
		fmt.Print(i, ",")
	}

	ss1 := "aあ"
	fmt.Println(ss1[0])
	fmt.Println(ss1[1], ss1[2], ss1[3])
	fmt.Println(ss1[1:4])

	b1 := make([]byte, 0, 5)
	b1 = append(b1, "abcde"...)
	fmt.Println(string(b1))

	copy(b1, "123")
	fmt.Println(string(b1))

	const size = 30
	const c = "a"
	b2 := make([]byte, 0, size)
	for i := 0; i < size; i++ {
		b2 = append(b2, c...)
	}
	fmt.Println(string(b2))

	s11 := ""
	for i := 0; i < size; i++ {
		s11 += c
	}
	fmt.Println(s11)
}
