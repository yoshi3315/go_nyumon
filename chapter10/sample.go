package main

import "fmt"

type myInt int
type myStruct struct {
	x, y int
}

func (m myInt) add(n int) myInt {
	return m + myInt(n)
}

func (m myStruct) add(n int) myStruct {
	return myStruct{m.x + n, m.y + n}
}

func (m myInt) addValue(n myInt) {
	m += n
}

func (m *myInt) addPointer(n myInt) {
	*m += n
}

func (m myInt) methodValue() {
	fmt.Println("値レシーバ: ", m)
}

func (m *myInt) methodPointer() {
	if m == nil {
		fmt.Println("ポインタレシーバ: nil")
	} else {
		fmt.Println("ポインタレシーバ: ", *m)
	}
}

func (m myInt) method(s string) {
	fmt.Println(s, m)
}

type point2d struct {
	x, y int
}

func (p2d point2d) mulXY() int {
	return p2d.x * p2d.y
}

func (p2d point2d) mul() int {
	return p2d.x * p2d.y
}

type point3d struct {
	point2d
	z int
}

func (p3d point3d) mul() int {
	return p3d.x * p3d.y * p3d.z
}

func main() {
	var v myInt = 1
	fmt.Println(v.add(2))

	var s = myStruct{4, 5}
	fmt.Println(s.add(6))

	v.addValue(2)
	fmt.Println(v)

	var p *myInt = &v
	p.addPointer(2)
	fmt.Println(*p)

	var v1 myInt = 1
	var p1 *myInt = &v1
	p1.methodValue()
	v1.methodPointer()

	p1 = nil
	p1.methodPointer()
	// p1.methodValue()

	var v2 myInt = 1
	methodValue1 := v2.method
	methodValue1("メソッド値:")

	methodExpr := myInt.method
	methodExpr(2, "値のメソッド式:")

	methodExprPtr := (*myInt).method
	v2 = 3
	methodExprPtr(&v, "ポインタのメソッド式:")

	p3d := point3d{point2d{2, 3}, 4}
	fmt.Println(p3d.mulXY())
	fmt.Println(p3d.mul())
	fmt.Println(p3d.point2d.mul())
}
