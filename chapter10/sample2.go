package main

import "fmt"

type myIF interface {
	typeName() string
	add(n int) int
}

type myInt int

func (m myInt) typeName() string {
	return "myInt"
}

func (m myInt) add(n int) int {
	return int(m) + n
}

func (m myInt) sub(n int) int {
	return int(m) - n
}

type myString string

func (m myString) typeName() string {
	return "myString"
}

func (m myString) add(n int) int {
	return len(m) + n
}

type embedIF interface {
	method1()
}

type myIF1 interface {
	embedIF
	method2()
}

func (i myInt) method1() {
	fmt.Println("method1: ", i)
}

func (i myInt) method2() {
	fmt.Println("method2: ", i)
}

func typeAssertion(i interface{}) {
	n := i.(int)
	fmt.Println(n + 1)
}

func typeAssertionCheck(i interface{}) {
	n, ok := i.(int)
	if ok {
		fmt.Println("int: ", n+1)
	}

	s, ok := i.(string)
	if ok {
		fmt.Println("string: ", s)
	}
}

func typeSwitch(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int: ", v)
	case string:
		fmt.Println("string: ", v)
	case int32, int64:
		fmt.Println("int32/64: ", v)
	default:
		fmt.Println("default :", v)
	}
}

func typeSwitchNil(i interface{}) {
	switch v := i.(type) {
	case nil:
		fmt.Println("nil: ", v)
	case *int:
		fmt.Println("*int: ", v)
	}
}

func main() {
	var n myInt = 1
	var i myIF = n
	fmt.Println(i.typeName())
	fmt.Println(i.add(2))

	i = &n
	fmt.Println(i.typeName())
	fmt.Println(i.add(2))

	var s myString = "abc"
	i = s
	fmt.Println(i.typeName())
	fmt.Println(i.add(2))

	var i0 myIF1 = myInt(1)
	i0.method1()
	i0.method2()

	var i1 interface{} = int(1)
	var i2 interface{} = byte(1)

	fmt.Println("i1 == i2: ", i1 == i2)

	var i3 interface{} = nil
	var i4 interface{} = (*int)(nil)

	fmt.Println("i3 == i4: ", i3 == i4)

	fmt.Println("i3 == nil: ", i3 == nil)
	fmt.Println("i4 == nil: ", i4 == nil)

	fmt.Println("i3 == (*int)(nil): ", i3 == (*int)(nil))
	fmt.Println("i4 == (*int)(nil): ", i4 == (*int)(nil))

	typeAssertion(1)
	// typeAssertion("a")

	typeAssertionCheck(1)
	typeAssertionCheck("a")

	typeSwitch(1)
	typeSwitch("a")
	typeSwitch(int32(1))
	typeSwitch(1.1)

	typeSwitchNil(nil)
	var nn *int = nil
	typeSwitchNil(nn)
}
