package main

import (
	"encoding/json"
	"fmt"
)

type JsonData struct {
	A string
	B int
}

func unmarshal(jsonText string) {
	var data JsonData
	err := json.Unmarshal(([]byte)(jsonText), &data)

	if syntaxErr, ok := err.(*json.SyntaxError); ok {
		fmt.Println("[文法エラー]", syntaxErr)
		fmt.Println(" Offset:", syntaxErr.Offset)
		return
	}

	if typeErr, ok := err.(*json.UnmarshalTypeError); ok {
		fmt.Println("型エラー：", typeErr)
		fmt.Println(" Offset:", typeErr.Offset)
		fmt.Println(" Value:", typeErr.Value)
		fmt.Println(" Type:", typeErr.Type)
		return
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("変換OK")
	fmt.Println(" ", data)
}

func main() {
	// i, err := strconv.Atoi("10")
	// if err != nil {
	// 	fmt.Println("エラー出力", err)
	// 	return
	// }
	// fmt.Println(i + 1)

	// i, err = strconv.Atoi("あ")
	// if err != nil {
	// 	fmt.Println("エラー出力", err)
	// 	return
	// }
	// fmt.Println(i + 1)

	unmarshal(`{"A": "X", "B": 1}`)
	unmarshal(`{"A": "X", "B": 1}`)
	unmarshal(`{"A": 2, "B": 1}`)
}
