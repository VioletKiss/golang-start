package main

import "fmt"

func main() {

	// 声明切片，初始化3个长度的默认值，容量为 5，len = 3, [0,0,0], cap = 5
	var numbers = make([]int, 3, 5)

	// 向numbers追加一个元素1，此时numbers为 len = 4，[0,0,0,1],cap = 5
	numbers = append(numbers, 1)

	fmt.Printf("len = %d, slice = %v\n", len(numbers), numbers)

	// 向numbers追加一个元素2，此时numbers为 len = 5，[0,0,0,1,2],cap = 5
	numbers = append(numbers, 2)

	fmt.Printf("len = %d, slice = %v\n", len(numbers), numbers)

	// 向容量cap已满的numbers追加一个元素,此时numbers为 len = 5，[0,0,0,1,2,3], cap = 10
	// 扩容原数组的一倍
	numbers = append(numbers, 3)

	fmt.Printf("len = %d, slice = %v\n", len(numbers), numbers)

}
