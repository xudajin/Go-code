package main

import "fmt"

var r string = "哈哈"

func main()  {
	// +号的使用
	//当 int + int 时  进行加法运算
	var z int =10
	var x int =20 
	var c int = z + x
	fmt.Println("c=",c)
	// 当 string  + string 时，进行字符串拼接
	var q string = "我是"
	var w string = "你爹"
	var e string = q +w + r 
	fmt.Println("e=",e)

	// go 中的数据类型有哪些
	
//基本数据类型:数值型，浮点数类型，字符型(byte)，布尔型，字符串型
//复杂数据类型: 指针，数组，结构体，管道，函数，切片，接口，map(类似集合)

	 



	
}