package main

import "fmt"

// interface{}是万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)
	// interface{} 该如何区分 此时引用的底层数据类型到底是什么

	// 给 interface{} 提供“类型断言”的机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
		fmt.Printf("value type is %T\n", arg)
	} else {
		fmt.Println("arg is string type, value = ", value)
	}
}

type Tree struct {
	color string
}

func main() {
	tree := Tree{"green"}

	myFunc(tree)
	myFunc(100)
	myFunc("abc")
	myFunc(3.14)
}
