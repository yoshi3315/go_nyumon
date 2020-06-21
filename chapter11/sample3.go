package main

import (
	"fmt"
	"time"
)

const c = 100000

func funcMapLoop(m map[int]int) {
	for n := 0; n < c; n++ {
		m[n] = 1
	}
}

type myStruct struct {
	s string
	x int
}

func main() {
	map1 := map[string]int{"a": 1, "b": 2}

	map1["a"], map1["c"] = 3, 4
	fmt.Println("len(map1): ", len(map1))
	fmt.Println("map1: ", map1["a"], map1["b"], map1["c"], map1["d"])

	map2 := map1
	map1["c"] = 5
	fmt.Println("map2: ", map2["a"], map2["b"], map2["c"])

	var map3 map[string]int
	fmt.Println("map3: ", map3["a"])
	// map3["a"] = 1

	n, ok := map1["a"]
	fmt.Println(n, ok)

	n, ok = map1["e"]
	fmt.Println(n, ok)

	delete(map1, "a")
	n, ok = map1["a"]
	fmt.Println(n, ok)

	t1 := time.Now()
	map11 := make(map[int]int)
	funcMapLoop(map11)
	fmt.Println("map11: ", time.Since(t1))

	t2 := time.Now()
	map12 := make(map[int]int, c)
	funcMapLoop(map12)
	fmt.Println("map12: ", time.Since(t2))

	map1a := map[myStruct]int{
		{"A", 1}: 10,
		{"B", 2}: 20,
	}

	map2a := map[int]myStruct{
		10: {"A", 1},
		20: {"B", 2},
	}

	fmt.Println(map1a[myStruct{"A", 1}], map1a[myStruct{"B", 2}])
	fmt.Println(map2a[10], map2a[20])

	map1b := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

	for k, e := range map1b {
		fmt.Print(k, ":", e, ", ")
	}
	fmt.Println()

	for k, e := range map1b {
		fmt.Print(k, ":", e, ", ")
	}
	fmt.Println()
}
