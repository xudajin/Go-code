package main

import (
	"fmt"
)
// float类型的使用
func main()  {
	// 浮点数是有符号的(有负值)，浮点数有固定的范围和字段长度
	var price float32 = 32.25
	fmt.Println(price)

	var num2 float32 = -252.65212
	fmt.Println(num2)
	// float64 的精度 高于float32

	//Golang的浮点型默认声明为float64 类型
	var num5 =1.1 
	// Printf() 变量的数值类型
	fmt.Printf("num5的数据类型是%T\n",num5)

	// 科学计数法形式
	var num8 = 5.1234e2  //5.1234 * 10的2次方 e2 
	fmt.Println("num8=",num8)

	// 通常情况下推荐使用float64位



	
}

