package lib1

import "fmt"

func Lib2Test() {
	fmt.Println('Lib2Test')
}

func init() {
	fmt.Println("lib2.init()")
}
