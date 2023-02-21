package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1

	for {
		select {
		case c <- y:
			x, y = y, x+y
			fmt.Println("send")
		case <-quit:
			fmt.Println("quit")
			fmt.Println("x=", x, " y=", y)
			return
		}
	}
}

func main() {
	// 定义一个带有缓冲的channel
	c := make(chan int)
	quit := make(chan int)

	go func() {
		defer fmt.Println("goroutine 结束")
		fmt.Println("goroutine 正在运行...")
		time.Sleep(5 * 1000)
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
	fmt.Println("main goroutine 结束")
}
