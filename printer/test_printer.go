package main

import "fmt"

func changeValue(p int) {
	p = 10
}

// 一级指针
func printerChangeValue(p *int) {
	*p = 10
}

func main() {
	var a int = 1
	changeValue(a)
	fmt.Println("changeValue-a:", a)

	printerChangeValue(&a)

	fmt.Println("printChangeValue-a:", a)

	var p *int
	p = &a
	fmt.Println("p:", p)
	fmt.Println("&a:", &a)

	// 二级指针
	var pp **int
	pp = &p

	fmt.Println("&p", &p)
	fmt.Println("pp", pp)
}
