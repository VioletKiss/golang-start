package main

import "fmt"

func main() {
	// 声明切片并初始化，长度3
	//slice1 := []int{1,2,3}

	// 声明切片，但没有分配空间
	var slice1 []int
	//slice1 = make([]int, 3) // 开辟3个空间，默认值是0

	// 声明切片，同时分配3个空间，初始化为默认值
	//var slice1 []int = make([]int, 3)

	// 声明切片，同时分配3个空间，初始化为默认值
	//slice1 := make([]int, 3)

	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)

	// 判断一个slice是否为0
	if slice1 == nil {
		fmt.Println("slice1 是一个空切片")
	} else {
		fmt.Println("slice1 是有空间的")
	}

}
