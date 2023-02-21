package main

import "fmt"

// 如果固定4个元素的数组那么传进来的数组必须是[4]int的类型，传[10]int失败
// 函数内修改影响原数组（引用传递）
func printSlice(myArray []int) {
	// _ 表示匿名的变量
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
	myArray[0] = 100
}

func main() {
	// 固定长度的数组
	myArray := []int{1, 2, 3, 4} // 动态数组，切片 slice

	fmt.Printf("myArray type is %T\n", myArray)

	printSlice(myArray)

	fmt.Println("==========")

	for _, value := range myArray {
		fmt.Println("value = ", value)
	}

}
