package main

import (
	"fmt"
	"time"
)

func main() {
	// 定义一个带有缓冲的channel
	c := make(chan int, 3)
	fmt.Println("len(c) = ", len(c), ", cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("goroutine 结束")
		fmt.Println("goroutine 正在运行...")
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("goroutine 正在运行，发送的元素=", i, "， len(c) = ", len(c), ", cap(c) = ", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c // 从c中接受数据并赋值给num
		fmt.Println("num = ", num)
		time.Sleep(2 * time.Second)
	}

	fmt.Println("main goroutine 结束")
}
