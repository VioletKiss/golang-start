package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book struct {
}

func (book *Book) ReadBook() {
	fmt.Println("Read a Book")
}

func (book *Book) WriteBook() {
	fmt.Println("Write a Book")
}

func main() {
	// b: pair<type:Book,value:book{}地址>
	b := &Book{}

	// b: pair<type:Book,value:book{}地址>
	var r Reader
	r = b
	r.ReadBook()
	// b: pair<type:Book,value:book{}地址>
	var w Writer
	w = r.(Writer) // 此处的断言为什么会成功？因为w r 具体的type是一致的
	w.WriteBook()

}
