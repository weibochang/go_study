package lib1

import "fmt"

// 函数名称是大写，表示该函数可以被外部的包调用
func Lib1Test() {
	fmt.Println('Lib1Test')
}

func init() {
	fmt.Println("lib1.init()")
}
