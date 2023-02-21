//go:build ignore
// +build ignore

package main

import "fmt"

// const 可以定义枚举
const (
	BEIJING  = 0
	SHANGHAI = 1
	SHENZHEN = 2
)

// iota 关键字
const (
	// 第一行的iota的默认值是0，每行的iota都会累加1
	APPLE = iota
	BANANA
	CHERRY
	PEAR
)

const (
	a, b = iota + 1, iota + 2
	c, d
	e, f

	g, h = iota * 2, iota * 3
	i, k
	l, m = 100, 200
	n, o
)

func main() {

	// 常量 （只读，不可修改）
	const length int = 1

	fmt.Println("length = ", length)

	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("SHENZHEN = ", SHENZHEN)

	fmt.Println("APPLE = ", APPLE)
	fmt.Println("BANANA = ", BANANA)
	fmt.Println("CHERRY = ", CHERRY)
	fmt.Println("PEAR = ", PEAR)

	fmt.Println("a = ", a, "b = ", b)
	fmt.Println("c = ", c, "d = ", d)
	fmt.Println("e = ", e, "f = ", f)
	fmt.Println("g = ", g, "h = ", h)
	fmt.Println("i = ", i, "k = ", k)
	fmt.Println("l = ", l, "m = ", m)
	fmt.Println("n = ", n, "o = ", o)

}
