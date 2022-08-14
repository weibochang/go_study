package main

import "fmt"

func main() {

	// 声明一个slice切片，并且初始化，默认值是1，2，3，长度是3
	//slice1 := []int{1, 2, 3}

	// 声明slice1是一个切片，但是没有给分配空间
	//var slice1 []int
	//slice1 = make([]int, 3) // 开辟3个空间
	//slice1[0] = 100

	// 声明slice1是一个切片，同时给slice分配空间，3个空间，初始值是0
	//var slice1 = make([]int, 3)

	// 声明slice1是一个切片，同时给slice分配空间，3个空间，初始值是0，通过:=推到出来slice是一个切片
	slice1 := make([]int, 3)

	// %v 表示打印出全部的详细信息
	fmt.Printf("len = %d , slice = %v\n", len(slice1), slice1)

	//判断一个slice是否为空
	if slice1 == nil {
		fmt.Println("slice 是一个空切片")
	} else {
		fmt.Println("slice 是有空间的")
	}
}
