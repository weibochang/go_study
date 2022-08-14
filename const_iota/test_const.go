package main

import (
	"fmt"
)

func main() {
	// const 声明常量，（一般不能被修改的变量被称为常量）（只读）
	const aa = 1
	const (
		x = 1
		y = 2
	)
	const (
		// const 中可以使用关键字 ista，默认从0开始，之后每行累加1
		// iota 只能够配合const() 一起使用， iota只有在const进行累加效果。
		beijin = iota
		tianjin
		shanghai
	)

	const (
		a, b = iota + 1, iota + 2 // iota = 0, a = iota + 1, b = iota + 2, a = 1, b = 2
		c, d                      // iota = 1, c = iota + 1, d = iota + 2, c = 2, d = 3
		e, f                      // iota = 2, e = iota + 1, f = iota + 2, e = 3, f = 4

		g, h = iota * 2, iota * 3 // iota = 3, g = iota * 2, h = iota * 3, g = 6, h = 9
		i, k                      // iota = 4, i = iota * 2, k = iota * 3 , i = 8, k = 12
	)
	fmt.Println("this is aa:", aa)
	fmt.Println("this is x:", x, ",this is y:", y)

	fmt.Println("beijin:", beijin, ",tianjin:", tianjin, ",shanghai:", shanghai)

	fmt.Println("a = ", a, "b = ", b)
	fmt.Println("c = ", c, "d = ", d)
	fmt.Println("e = ", e, "f = ", f)

	fmt.Println("g = ", g, "h = ", h)
	fmt.Println("i = ", i, "k = ", k)

}
