package main

import (
	"fmt"
	"unsafe"
)
//整数类型的基本介绍
func main()  {
	// 有符号的整数 int8   范围-128 ~ 127  超出范围报错
	
	var j int8 = -128
	fmt.Println(j)

	//无符号的整数 uint8 没有负值 范围从 0~255
	var i uint8  = 255
	fmt.Println(i)

	// 整数的使用细节
	// int 为有符号 uint为无符号
	
	var n1 int64 =100  // 默认为 int 类型
	// 如何查看变量的数据类型  fmt.Printf() 用于格式化输出
	fmt.Printf("n1的类型为%T \n",n1) // %T 表示 type(类型)

	// 如何在程序中查看变量的字节大小和数据类型
	fmt.Printf("n2的类型%T  n2占用的字节数是%d\n", n1,unsafe.Sizeof(n1))

	// 在程序正确运行的前提下，尽量使用占用空间小的数据类型
	var age int8 =23  
	fmt.Println(age)
	// 一个byte = 8bit

}