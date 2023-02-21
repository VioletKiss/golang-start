package main

import "fmt"

// 接口（本质是一个指针）
type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

// 具体的类
type Cat struct {
	color string
}

func (cat *Cat) GetColor() string {
	return cat.color
}

func (cat *Cat) GetType() string {
	return "Cat"
}

func (cat *Cat) Sleep() {
	fmt.Println("Cat is sleep")
}

type Dog struct {
	color string
}

func (dog *Dog) Sleep() {
	fmt.Println("Dog is sleep")
}

func (dog *Dog) GetColor() string {
	return dog.color
}

func (dog *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal AnimalIF) {
	animal.Sleep()
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("kind = ", animal.GetType())
}

func main() {
	var animal AnimalIF
	animal = &Cat{"Green"}
	animal.Sleep()

	animal = &Dog{"Yellow"}
	animal.Sleep()

	cat := Cat{"Green"}
	dog := Dog{"Yellow"}

	showAnimal(&cat)
	showAnimal(&dog)
}
