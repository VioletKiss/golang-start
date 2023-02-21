package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (user User) Call() {
	fmt.Println("user is called...")
	fmt.Printf("%v\n", user)
}

func DoFieldAndMethod(input interface{}) {
	// 获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is : ", inputType)
	// 获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is : ", inputValue)

	// 通过type获取里面的字段
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface() //获取字段对应的value
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 通过type获取里面的方法，调用
	// 注意 func 《(user User)》 Call() 里《》标记的地方，这里类型需要对应
	for i := 0; i < inputType.NumMethod(); i++ {
		method := inputType.Method(i)
		fmt.Printf("%s: %v\n", method.Name, method.Type)
	}
	inputType.Elem()
}

func main() {
	user := User{1, "Amy", 18}
	DoFieldAndMethod(user)
}
