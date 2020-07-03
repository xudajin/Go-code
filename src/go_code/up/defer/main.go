package main

import (
	"fmt"
)

func sum(n1 int, n2 int) int {

	//当执行到defer时，暂时不执行，会将defer语句压入到独立的栈中
	// 当函数执行完毕后再从defer存在的栈中，按先入后出的
	//方式出栈，后执行 (既先defer的后输出)
	defer fmt.Println("ok1 n1=",n1)
	defer fmt.Println("ok2 n2=", n2)
	n1++  // n1++ =11
	n2++  // n2++ =21
	res := n1 + n2
	fmt.Println("res=",res)
	return res
} 

/*
name := "DD"  等价于 var name string , name = "DD" 两个语句  
报错 赋值语句不能再函数体外执行

*/ 
func main()  {
// 函数中的defer
// 为什么需要defer ：程序中创建资源后在函数执行完毕后
// 及时释放资源

	res := sum(10,20)
	fmt.Println("===",res)
	// 打印的结果
	// res =30
	// ok2 n2= 20
	// ok1 n1= 10
	// === 30

	// defer 的细节说明
	// 当执行到defer语句时，先不执行，存入栈中（先入后出）
	// 当将defer语句放入栈中，也会将相关的值拷贝后，放入栈中
	// 当后续的操作不会影响先defer 入栈的值
	// defer 最主要的价值是在，当函数执行完毕后，可以及时
	// 释放函数创建的资源
	// defer语句 是在整个函数执行结束后再执行的，所以
	// 可以先执行defer 不会影响函数的执行
	// 这中机制，使得程序员不用在为什么时候关闭资源而烦心





}