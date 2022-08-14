/*
 知识点一: defer 和 return 谁先谁后
	如果 return 和 defer 同时存在一个函数中，return要先执行，defer后执行
*/
package main

import "fmt"

func deferFunc() int {
	fmt.Println("defer func called")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called")
	return 0
}

func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	returnAndDefer()
}
