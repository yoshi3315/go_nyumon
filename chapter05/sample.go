package main

import "fmt"

var (
	v1 = 2 + 3

	u1 uint8 = 255
	v2       = u1 + 1

	u2 int8 = 127
	v3      = u2 + 1

	u3 uint8 = 0xA3
	u4 uint8 = 0x53

	v4 = u3 & u4
	v5 = u3 | u4
	v6 = u3 ^ u4
	v7 = u3 &^ u4
	v8 = ^u3

	v9  = u3 << 4
	v10 = u3 >> 4

	i1  int8 = -0x50
	v11      = i1 << 4
	v12      = i1 >> 4
)

func main() {
	fmt.Println(v1, v2, v3)
	fmt.Println(v4, v5, v6, v7, v8)
	fmt.Println(v9, v10)
	fmt.Println(v11, v12)

	v1 += 1
	fmt.Println(v1)

	fmt.Println(1 == 1, 1 != 2)
	fmt.Println(1 <= 2, 1 < 2)
	fmt.Println(2 >= 1, 2 > 1)

	fmt.Println("ABC" == "ABC", "A" < "B", "ABC" < "ACB")

	fmt.Println(trueFunc() || falseFunc())
	fmt.Println(falseFunc() && trueFunc())
	fmt.Println(falseFunc() || falseFunc())
	fmt.Println(trueFunc() && trueFunc())

	i := 1
	i++
	fmt.Println(i)

	i--
	fmt.Println(i)
}

func trueFunc() bool {
	fmt.Print("[trueFunc]")
	return true
}

func falseFunc() bool {
	fmt.Print("[falseFunc]")
	return false
}
