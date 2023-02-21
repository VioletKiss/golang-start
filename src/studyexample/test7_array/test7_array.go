package main

import "fmt"

// 如果固定4个元素的数组那么传进来的数组必须是[4]int的类型，传[10]int失败
// 函数内修改数组值不影响原数组(值传递)
func printArray(myArray [4]int) {
	for index, value := range myArray {
		fmt.Println("index = ", index, ", value = ", value)
	}
	myArray[0] = 100
}

func main() {
	// 固定长度的数组
	var myArray1 [10]int

	myArray2 := [10]int{1, 2, 3, 4}

	myArray3 := [4]int{1, 2, 3, 4}

	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}

	for index, value := range myArray2 {
		fmt.Println("index = ", index, ", value = ", value)
	}

	printArray(myArray3)

	for index, value := range myArray3 {
		fmt.Println("index = ", index, ", value = ", value)
	}

}
