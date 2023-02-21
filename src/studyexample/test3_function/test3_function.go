//go:build ignore
// +build ignore

package main

import "fmt"

func func1(var1 string, var2 int) int {
	fmt.Println("-------------func1-------------")
	fmt.Println("var1 = ", var1)
	fmt.Println("var2 = ", var2)

	c := 100
	return c
}

// 返回多个返回值，匿名
func func2(var1 string, var2 int) (int, int) {
	fmt.Println("-------------func1-------------")
	fmt.Println("var1 = ", var1)
	fmt.Println("var2 = ", var2)

	return 666, 777
}

// 返回多个返回值，有形参名称(没有赋值将返回默认值)
func func3(var1 string, var2 int) (r1 int, r2 int) {
	fmt.Println("-------------func1-------------")
	fmt.Println("r1 = ", r1)
	fmt.Println("r2 = ", r2)
	// 给形参赋值
	r1 = 1000
	r2 = 2000

	return
}

func main() {
	r1 := func1("fun1", 1)
	fmt.Println("r1 = ", r1)

	r1, r2 := func2("func2", 1)
	fmt.Println("r1 = ", r1, "r2 = ", r2)

	r1, r2 = func3("func3", 1)
	fmt.Println("r1 = ", r1, "r2 = ", r2)

}
