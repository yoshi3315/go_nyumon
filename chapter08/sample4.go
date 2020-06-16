package main

import "fmt"

func funcArg() string {
	fmt.Println("funcArg関数の実行")
	return "funcArgの戻り値"
}

func funcDefer1() {
	defer fmt.Println(funcArg())
	fmt.Println("funcDefer1関数終了")
}

func funcDefer2() {
	funcVar := func() { fmt.Println("defer文登録関数の実行") }

	defer funcVar()

	funcVar = func() { fmt.Println("使用されない") }
}

func funcDeferResult() (result int) {
	defer func() {
		fmt.Println("更新前戻り値: ", result)
		result += 2
	}()

	return 1
}

func main() {
	fmt.Println("main開始")

	defer fmt.Println("defer呼び出し")

	fmt.Println("main終了")

	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")

	funcDefer1()
	funcDefer2()

	fmt.Println("関数戻り値: ", funcDeferResult())
}
