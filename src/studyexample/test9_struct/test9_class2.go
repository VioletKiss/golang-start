package main

import "fmt"

// 如果类名称首字母大写，表示其他包也能够访问（同理属性和方法）
type Human struct {
	name string
	sex  string
}

func (human *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (human *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

type SuperMan struct {
	Human // 继承Human类的方法
	level int
}

// 重定义父类方法
func (superMan *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

// 子类新方法
func (superMan *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func (superMan *SuperMan) print() {
	fmt.Printf("name = %s,sex = %s,level = %d\n", superMan.name, superMan.sex, superMan.level)
}

func main() {
	human := Human{name: "zhang3", sex: "female"}
	human.Eat()
	human.Walk()

	superMan := SuperMan{Human{"li4", "female"}, 99}
	//var s SuperMan  //定义对象
	//s.name = "li4"
	//s.sex = "male"
	//s.level = 88

	superMan.Eat()
	superMan.Walk()
	superMan.Fly()

	superMan.print()

}
