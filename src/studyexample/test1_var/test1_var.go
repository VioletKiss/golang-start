package main

/*
	四种变量的声明方式
*/
import "fmt"

// 声明全局变量，只有123是可以的
var gA int = 1
var gB = 1

func main() {
	// 1. 声明一个变量，默认值为0
	var a int
	fmt.Println("a = ", a)

	// 2.声明一个变量，初始化一个值
	var b int = 1
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	// 3.在初始化是时可以省去数据类型
	var c = 1
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	// 4. 省去var关键字，自动匹配
	e := 1
	fmt.Println("e = ", e)

	// 打印全局变量
	fmt.Println("gA = ", gA)
	fmt.Println("gB = ", gB)

	// 声明多个变量
	var xx, yy int = 1, 2
	fmt.Println("xx = ", xx)
	fmt.Println("yy = ", yy)
	var kk, ll = 100, "string"
	fmt.Println("kk = ", kk)
	fmt.Println("ll = ", ll)

	// 多行的多变量声明
	var (
		vv int  = 1
		jj bool = true
	)
	fmt.Println("vv = ", vv)
	fmt.Println("jj = ", jj)
}
