package main



import (
	"fmt"
)
//golang中的指针类型
func main()  {
	//基本数据类型的内存布局
	var i int = 10 
	// i 内存地址是什么
	fmt.Println("i的地址为",&i) //&i 打印i的地址
	
	// var ptr *int = &i  // &i 为i空间地址
	// ptr 是一个指针变量,指向一个空间，空间存放的也是空间地址
	// ptr 的类型是*int
	// ptr 本身的值是&i
	// var ptr *int = &i
	// fmt.Printf("ptr=%v\n",ptr)
	// fmt.Printf("ptr的地址=%v\n",&ptr)
	// fmt.Printf("ptr 指向的值=%v",*ptr)

	var ptr *int  // 声明一个指针变量
	ptr = &i     // i的内存地址 赋值给指针变量ptr，ptr是空间地址
	*ptr = 20    // *符号表示使用ptr所指向的地址去获取i的值
	fmt.Println("i的值是",i)


	//指针的说明
	// 1 值类型，都有对应的指针类型，形式为 *数据类型
	// 如 int对应的指针就是 *int ，*float32 对应 *float32

	// 2 值类型包括 : 基本数据类型 int系列 ，float 系列 bool string
	// 数组 和 结构体

	 // 值类型 和引用类型
	 // 1 值类型 : 基本数据类型 int系列 ，float 系列 bool string
	// 数组 和 结构体
	// 2 引用类型 : 切片 ，map ，管道 ，interface ，指针









	
}