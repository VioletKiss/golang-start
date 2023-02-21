package main

import "fmt"

func main() {
	//1.声明map类型，key是string，value是string
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1 is nil")
	}

	// 在使用map前，需要先给map分配空间
	myMap1 = make(map[string]string, 10)

	myMap1["one"] = "java"
	myMap1["two"] = "go"

	fmt.Println(myMap1)

	//2.声明map
	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "go"

	fmt.Println(myMap2)

	//3.声明map
	myMap3 := map[string]string{
		"one": "java",
		"two": "go",
	}

	fmt.Println(myMap3)

	// map的使用

	cityMap := make(map[string]string)
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	//遍历
	printMap(cityMap)

	//删除
	delete(cityMap, "China")

	// 修改
	cityMap["USA"] = "DC"

	fmt.Println("=========")

	//遍历
	printMap(cityMap)

}

func printMap(cityMap map[string]string) {
	//cityMap是引用传递
	//遍历
	for key, value := range cityMap {
		fmt.Println("key = ", key, " value = ", value)
	}
}
