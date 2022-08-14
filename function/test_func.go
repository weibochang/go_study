package main

import "fmt"

// 多个参数，单个返回值（匿名）
func foo1(a string, b int) int {
	fmt.Printf("this is a:%s, type is %T ; this is b:%d type is %T\n", a, a, b, b)
	c := b + 10
	return c
}

// 多个参数，多个返回值（匿名）
func foo2(a string, b int) (int, int) {
	fmt.Printf("this is a:%s, type is %T ; this is b:%d type is %T\n", a, a, b, b)
	return 3, 4
}

// 多个参数，多个返回值（有形参名称）
func foo3(a string, b int) (r1 int, r2 int) {
	// r1 r2 是形参，默认值是0
	fmt.Println("this is r1:", r1, ",this is r2:", r2)

	fmt.Printf("this is a:%s, type is %T ; this is b:%d type is %T\n", a, a, b, b)

	r1 = 1000
	r2 = 2000

	return r1, r2
}

// 多个参数，多个返回值（有形参名称）
func foo4(a string, b int) (r1 int, r2 int) {
	// r1 r2 是形参，默认值是0
	fmt.Println("this is r1:", r1, ",this is r2:", r2)

	fmt.Printf("this is a:%s, type is %T ; this is b:%d type is %T\n", a, a, b, b)

	r1 = 1000
	r2 = 2000

	// 定义形参之后，直接return 就会把形参都给返回
	return
}

func main() {
	func_foo1 := foo1("abc", 1)
	fmt.Println("this is func_foo1:", func_foo1)

	func_foo2_1, func_foo2_2 := foo2("abc", 1)
	fmt.Println("this is func_foo2_1:", func_foo2_1, "; this is func_foo2_2:", func_foo2_2)

	func_foo3_1, func_foo3_2 := foo3("abc", 1)
	fmt.Println("this is func_foo3_1:", func_foo3_1, "; this is func_foo3_2:", func_foo3_2)

	func_foo4_1, func_foo4_2 := foo4("abc", 1)
	fmt.Println("this is func_foo4_1:", func_foo4_1, "; this is func_foo4_2:", func_foo4_2)
}
