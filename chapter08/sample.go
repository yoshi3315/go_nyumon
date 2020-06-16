package main

import "fmt"

func funcName1() {
	fmt.Println("Hello")
}

func funcParam1(n int) {
	fmt.Println(n)
}

func funcParam2(n1, n2 int) {
	fmt.Println(n1 + n2)
}

func funcVariadic1(slice ...string) {
	for _, s := range slice {
		fmt.Print(s, ", ")
	}
	fmt.Println()
}

func funcVariadic2(n int, slice ...string) {
	fmt.Print(n, ": ")

	for _, s := range slice {
		fmt.Print(s, ", ")
	}
	fmt.Println()
}

func funcResult1() int {
	return 1
}

func funcResult2() (int, int) {
	return 1, 2
}

func funcResult3(x, y int) {
	fmt.Println(x + y)
}

func funcNamedResult1() (i int) {
	i = 10
	return
}

func funcNamedResult2() (i int) {
	i = 10
	return 20
}

func funcNamedResult3() (i int, j int) {
	i, j = 10, 20
	return
}

func funcNamedResult4() (i, j int) {
	i, j = 10, 20
	return
}

func funcNamedResult5() (i int) {
	return
}

func main() {
	funcName1()

	funcName2 := funcName1

	funcName2()

	funcParam1(1)
	funcParam2(1, 2)

	funcVariadic1()
	funcVariadic1("a")
	funcVariadic1("a", "b")

	slice := []string{"a", "b"}
	funcVariadic1(slice...)

	funcVariadic2(10, "a", "b")

	a := funcResult1()
	fmt.Println(a)

	a, b := funcResult2()
	fmt.Println(a, b)

	funcResult3(funcResult2())

	fmt.Println(funcNamedResult1())
	fmt.Println(funcNamedResult2())
	fmt.Println(funcNamedResult3())
	fmt.Println(funcNamedResult4())
	fmt.Println(funcNamedResult5())
}
