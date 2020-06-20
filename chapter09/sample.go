package main

import (
	"fmt"
	"reflect"
)

type myInt32 int32

type (
	myInt    int
	myString string
)

type point struct {
	name string
	x, y int64
}

type point3 struct {
	point
	y, z int64
}

type point4 struct {
	name string `color:"red" size:"12pt"`
	x, y int64
}

func main() {
	var n1 myInt32 = 1
	var n2 myInt32 = 2
	fmt.Println(n1 + n2)

	var p1 point
	p1.name = "point"
	p1.x = 1
	p1.y = 2

	fmt.Println("p1: ", p1.name, p1.x, p1.y)

	p2 := p1
	p1.name = "この設定はp2に反映されない"
	fmt.Println("p2: ", p2.name, p2.x, p2.y)

	p3 := point{"point", 1, 2}
	fmt.Println("p3: ", p3.name, p3.x, p3.y)

	p4 := point{name: "point", y: 2}
	fmt.Println("p4: ", p4.name, p4.x, p4.y)

	p5 := point{}
	fmt.Println("p5: ", p5.name, p5.x, p5.y)

	var p6 point3
	p6.name = "p1name"
	p6.x = 1
	p6.y = 2
	p6.point.y = 20
	p6.z = 3

	fmt.Println("name: ", p6.name, p6.point.name)
	fmt.Println("x: ", p6.x, p6.point.x)
	fmt.Println("y: ", p6.y, p6.point.y)
	fmt.Println("z: ", p6.z)

	var p7 point4
	field := reflect.TypeOf(p7).Field(0)
	fmt.Println("color: ", field.Tag.Get("color"))
	fmt.Println("size: ", field.Tag.Get("size"))
}
