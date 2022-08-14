package main

import "fmt"

// 这种方式 传递数组类型，参数长度一定要相等，否则会报错
func printArray(myArray [4]int) {
	// 值拷贝
	for index, value := range myArray {
		fmt.Println("index=", index, ",value=", value)
	}
	myArray[0] = 111
}

func main() {
	// 固定长度的数组
	var myArray [10]int

	// 指定了数组的前4个值
	myArray2 := [10]int{1, 2, 3, 4}

	myArray3 := [4]int{11, 22, 33, 44}

	for i := 0; i < len(myArray); i++ {
		fmt.Println(i, myArray[i])
	}
	fmt.Println("using range keyword")

	for index, value := range myArray {
		fmt.Println("index:", index, ",value:", value)
	}

	fmt.Println("test myArray2")
	for index, value := range myArray2 {
		fmt.Println("index:", index, ",value:", value)
	}

	fmt.Println("see data type")
	// 查看数据类型
	fmt.Printf("myArray1 type is %T\n", myArray)
	fmt.Printf("myArray2 type is %T\n", myArray2)

	fmt.Println("func use array params")
	printArray(myArray3)
	fmt.Println("myArray3:", myArray3)
}
