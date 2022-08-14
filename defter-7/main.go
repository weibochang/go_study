package main

// -- 基本内容 --
// defer 表示一个函数在 执行最后，结束之前 的一种机制
// defer 之后可以跟一个表达式，或者跟一个函数的调用
// defer 可以有多个
// 多个 defer 相当于栈的方式（后进先出）

import "fmt"

func main() {
	// 在 main 函数结束之前运行
	defer fmt.Println("main end -1")
	defer fmt.Println("main end -2")

	fmt.Println("main: hello go 1")
	fmt.Println("main: hello go 2")
}
