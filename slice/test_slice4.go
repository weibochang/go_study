// 切片容量的截取
// 默认截取之后，其实底层是指向同一个slice的，修改截取后的slice，原始的slice也会发生变化

package main

import "fmt"

func main() {
	s := []int{1, 2, 3}

	s1 := s[0:2]
	fmt.Println(s1)

	s1[0] = 100
	fmt.Println(s)
	fmt.Println(s1)

	// copy 可以将底层数组的slice一起进行拷贝
	s2 := make([]int, 3)
	copy(s2, s)
	fmt.Println(s2)

}
