package main


import (
	"fmt"
)


// 分支控制
//基本语法

func main()  {
	// 单分支
	// var age int 
	// fmt.Println("请输入你的年龄")
	// fmt.Scanln(&age)
	// if age > 18 {
	// 	fmt.Println("你已经成年")
		
	// }
	// go 直接在if语句中，直接定义变量,只能使用 num:= 10 推导式定义
	if num := 39 ; num > 10 {
		fmt.Println("num的zhi为",num)
	}

	//双分支
	var x int = 1
	if x < 5 {
		fmt.Println("小于")
	}else {
		fmt.Println("大于")
	}

	var d1 int = 10
	var d2 int = 10
	if d1 + d2 < 40 {
		fmt.Println("成立")
	} 
	// 多分支
	var q1 int =1
	if q1 <3 {
		fmt.Println("小于3")
	} else if q1 > 3 && q1 < 8 {
		fmt.Println("大于3，小于8")
	} else {
		fmt.Println("大于8")
	}







}