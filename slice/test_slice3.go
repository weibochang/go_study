// 切片容量的追加
package main

import "fmt"

func main() {

	var numbers = make([]int, 3, 5)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 向 numbers 追加一个元素 1，numbers len = 4, [0,0,0,1], cap = 5
	numbers = append(numbers, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 2)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// cap容量已经满了，如果新增元素，这时候会新增cap长度的容量
	numbers = append(numbers, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	fmt.Println("---------")
	var numbers2 = make([]int, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)

}
