package main


import "fmt"

// 匿名函数： 没有函数名的函数

// 全局匿名函数
var	Fun1 = func(n1 int ,n2 int) int {
		return n1 - n2 
	}



func main()  {
	//方式1:在定义时直接调用，该匿名函数只能调用一次
	res := func (n1 int ,n2 int) int { //定义函数
		return n1 + n2
	} (10,20)   // （10，20） 调用函数

	fmt.Println("res",res)


	//方式二: 将匿名函数func 赋值给 a
	a := func (n1 int , n2 int) int {
		return n1 -n2 	
	}
	res3 := a(30,10)
	fmt.Println("res3=",res3)


	res4 := Fun1(50,5)
	fmt.Println("res4=",res4)




}
