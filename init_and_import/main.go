package main

import (
	// 匿名导入 无法使用导入包的方法，但是可以执行包的 init() 方法
	//_ "go_response/init_and_import/lib1"

	// 给包起一个别名，叫 aa ，可以通过aa来调用包的方法
	//aa "go_response/init_and_import/lib1"

	// 导入包内所有的方法，可以直接使用包内的方法名称去调用Lib1Test，不需要使用 lib.Lib1Test
	//. "go_response/init_and_import/lib1"

	"go_response/init_and_import/lib1"
	"go_response/init_and_import/lib2"
)

func main() {
	lib1.Lib1Test()
	lib2.Lib2Test()
}
