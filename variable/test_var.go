package main

import "fmt"

func main() {
	// 1.声明一个变量，默认是0
	var a int

	// 2.声明一个变量，初始化一个值
	var b int = 100

	// 3.声明一个变化，不指定类型，可以通过值自动匹配当前数据的类型
	var c = 100

	// 4.(常⽤的⽅法) 省去var关键字，直接⾃动匹配(这种方法只能在函数内使用，不能用于全局变量的声明)
	d := 100

	// 多变量声明
	var aa, bb int = 1, 2
	var cc, dd = 3, 4
	var (
		x int     = 1
		y bool    = true
		z float32 = 3.1
	)

	fmt.Printf("this is a:%d,type is %T\n", a, a)
	fmt.Printf("this is b:%d,type is %T\n", b, b)
	fmt.Printf("this is c:%d,type is %T\n", c, c)
	fmt.Printf("this is d:%d,type is %T\n", d, c)

	fmt.Println("aa:", aa, "bb:", bb)
	fmt.Println("cc:", cc, "dd:", dd)
	fmt.Println("x:", x, "y:", y, "z:", z)
}
