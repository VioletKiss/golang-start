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
		for i := 0; i < 5; i++ {
			c <- i
			fmt.Println("goroutine 正在运行，发送的元素=", i, "， len(c) = ", len(c), ", cap(c) = ", cap(c))
		}
		// close可以关闭一个channel
		close(c)
	}()

	time.Sleep(2 * time.Second)

	// 必须关闭了channel死循环读取才不会报错不然会死锁

	//1
	for {
		// 从c中接受数据并赋值给num
		if num, ok := <-c; ok {
			fmt.Println("num = ", num)
			time.Sleep(2 * time.Second)
		} else {
			break
		}
	}

	//这里和1实现的效果一样，使用range更简洁
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("main goroutine 结束")
}
