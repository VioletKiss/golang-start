package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 2000, 10, []string{"zhouxingchi", "zhangbozhi"}}

	//编码 转为json格式
	jsonStr, err := json.Marshal(movie)

	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}
	fmt.Printf("jsonStr = %s\n", jsonStr)

	//解码 json转结构体
	//jsonStr = {"title":"喜剧之王","year":2000,"rmb":10,"actors":["zhouxingchi","zhangbozhi"]}
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal error", err)
		return
	}
	fmt.Printf("%v\n", myMovie)
}
