package main

import "fmt"

// 定义结构体
type Book struct {
	title string
	auth  string
}

func changeBook1(book Book) {
	// 传递一个book的副本
	book.auth = "111"
}

func changeBook2(book *Book) {
	// 指针传递
	book.auth = "111"
}

func main() {
	var book1 Book
	book1.title = "Golang"
	book1.auth = "zhang3"

	fmt.Printf("%v\n", book1)

	changeBook1(book1)

	fmt.Printf("%v\n", book1)

	changeBook2(&book1)

	fmt.Printf("%v\n", book1)

}
