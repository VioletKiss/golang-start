package main

import "fmt"

// 如果类名称首字母大写，表示其他包也能够访问（同理属性和方法）
type Hero struct {
	Name  string
	Ad    int
	Level int
}

func (hero *Hero) Show() {
	fmt.Println("Name = ", hero.Name)
	fmt.Println("Ad = ", hero.Ad)
	fmt.Println("Level = ", hero.Level)
}

func (hero *Hero) GetName() string {
	return hero.Name
}

func (hero *Hero) SetName(newName string) {
	hero.Name = newName
}

func main() {
	// 创建一个对象
	hero := Hero{Name: "zhang3", Ad: 100}
	hero.Show()
	hero.SetName("li4")
	hero.Show()
}
