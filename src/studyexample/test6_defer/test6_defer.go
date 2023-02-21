package main

import "fmt"

func deferFunc() {
	fmt.Println("defer func called ...")
}

func returnFunc() int {
	fmt.Println("return func called ...")
	return 0
}

func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	// 写入defer关键字，defer在函数周期结束后调用(先进后出)
	defer fmt.Println("main end1")
	defer fmt.Println("main end2")

	fmt.Println("main::hello go 1")
	fmt.Println("main::hello go 2")

	returnAndDefer()

}
