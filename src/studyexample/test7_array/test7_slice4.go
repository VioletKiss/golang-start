package main

import "fmt"

func main() {

	s := []int{1, 2, 3}

	// 左闭右开 从0（包含）到2（不包含），从0开始可以省略
	s1 := s[0:2]

	fmt.Println(s1)

	// 修改s1会影响s
	s1[0] = 100

	fmt.Println(s)

	fmt.Println(s1)

	// copy可以将底层数组一起拷贝
	s2 := make([]int, 3)

	// 将s中的值依次拷贝到s2中
	copy(s2, s)
	fmt.Println(s2)

}
