package main

import "fmt"

func main() {
	fmt.Println("main start...")
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
