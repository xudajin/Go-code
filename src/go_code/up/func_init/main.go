package main

import (
	"fmt"
	"go_code/up/func_init/utils"
)

//每一个源文件可以包含一个init函数，该函数会在main函数
//函数执行前，被GO框架调用

// inint 函数   (在main函数之前被调用)
//ps : 如果一个文件同时包含全局变量定义，init函数，main函数
// 则执行的流程是 全局变量定义 ——> init 函数 -> main

var age int = test()

func test() int {
	fmt.Println("test")
	return 90
	
}
// init函数
func init()  {
	fmt.Println("init")
}

func main()  {
	fmt.Println("main")
	fmt.Println("Age=",utils.Age,"Name=",utils.Name)

	// 如果main.go 和utils.go都含有变量定义，init函数
	//时，执行的流程是如何的
	/*
	1  先 执行执行utils中的变量定义
	2  执行utils 中的 init函数
	3  执行main.go 的变量定义
	4 执行main.go 的init函数
	5 执行main 函数
	ps  既 先执行被导入文件中的代码
	*/
	


}