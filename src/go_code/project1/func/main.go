package main

import (
	"fmt"
)

// 自定义函数
// 函数的定义
func cal(n1 float64, n2 float64, do byte) (float64) {
	var res float64
	switch do {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2 
	case '*': 
		res = n1 * n2

	default:
		fmt.Println("输入有误")
	}
	return res
}
//  函数的使用（重点）
func main()  {
	var result float64 =cal(1.2,3.2,'*') // 函数的调用
	fmt.Println(result)

	
	




	
}