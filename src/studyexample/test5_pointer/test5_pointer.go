//go:build ignore
// +build ignore

package main

import "fmt"

// 传入地址，给变量交换数值
func swap(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}

func main() {
	var a = 10
	var b = 20
	swap(&a, &b)

	fmt.Println("a = ", a, "b = ", b)

}
