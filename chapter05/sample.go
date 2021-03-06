package main

import (
	"fmt"
	"unsafe"
)

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

const (
	intLit       = 123
	strLit       = "ABC"
	intConst     = -intLit
	strConst     = strLit + "DEF"
	floatConst   = float64(intLit + 456)
	complexConst = complex(1.1, 2.2)
	realConst    = real(complexConst)
	imagConst    = imag(complexConst)
)

var a = [3]int{1, 2, 3}
var sss = struct {
	x int64
	y float64
}{}

const (
	lenConst      = len(a)
	capConst      = cap(a)
	alignofConst  = unsafe.Alignof(sss)
	offsetofConst = unsafe.Offsetof(sss.y)
	sizeofConst   = unsafe.Sizeof(sss)
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

	var (
		f32r float32 = 1.1
		f32i float32 = 2.2
		f64r float64 = 1.1
		f64i float64 = 2.2
	)

	var c64 complex64 = complex(f32r, f32i)
	fmt.Println("c64: ", c64)

	var c128 complex128 = complex(f64r, f64i)
	fmt.Println("c128: ", c128)

	f32r = real(c64)
	f32i = imag(c64)
	fmt.Println("f32r: ", f32r)
	fmt.Println("f32i: ", f32i)

	f64r = real(c128)
	f64i = imag(c128)
	fmt.Println("f64r: ", f64r)
	fmt.Println("f64i: ", f64i)

	s1 := "abc"
	s2 := "あいうえお"
	s3 := s1 + s2
	fmt.Println(s3)
	fmt.Println(len(s3))

	var (
		u8        uint8
		u16       uint16
		f64       float64
		r                = 'あ'
		str       string = "abcあいうえお"
		byteSlice []byte
		runeSlice []rune
	)

	u8 = uint8(u16)
	u16 = uint16(u8)

	f64 = float64(u16)
	u16 = uint16(f64)

	var ss1 string = string(r)
	fmt.Println(ss1)

	byteSlice = []byte(str)
	fmt.Println(byteSlice)

	runeSlice = []rune(str)
	fmt.Println(runeSlice)

	var ss2 string = string(byteSlice)
	fmt.Println(ss2)

	var ss3 string = string(runeSlice)
	fmt.Println(ss3)

	fmt.Println(intConst, strConst, floatConst)
	fmt.Println(complexConst, realConst, complexConst)
	fmt.Println(lenConst, capConst)
	fmt.Println(alignofConst, offsetofConst, sizeofConst)
}

func trueFunc() bool {
	fmt.Print("[trueFunc]")
	return true
}

func falseFunc() bool {
	fmt.Print("[falseFunc]")
	return false
}
