//go:build ignore
// +build ignore

package main

import (
	// 导包前使用_可以仅执行init方法
	_ "awesomeProject/src/studyexample/test4_init/lib1"
	mylib2 "awesomeProject/src/studyexample/test4_init/lib2"
	. "awesomeProject/src/studyexample/test4_init/lib3"
)

func main() {
	//lib1.Lib1Func()
	mylib2.Lib2Func()
	Lib3Func()
}
