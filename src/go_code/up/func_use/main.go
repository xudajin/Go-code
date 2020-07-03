package main

import "fmt"

func test(n1 int) (int) {
	n1 = n1 +1 
	fmt.Println("test n1=",n1)
	return n1
}
//         参数1   参数2     返回值
// test2(n1 int , n2 int) (int,int) {}
func test2(n1 int , n2 int) (int,int) {
	var dec int = n1 - n2
	var sum int = n1 + n2
	return dec,sum
}
// 函数可以作为参数，传入函数中
func myFun(fun func(int) int ,num1 int) int {
	return fun(num1)
}

func main()  {
	n1 := 10
	// 将 n1 作为参数 传入test中
	//调用函数时，会开辟新的栈区，函数执行结束时栈区会回收
	// 并执行该函数后的代码，函数的栈空间 和主函数的栈空间
	// 不同
	
	test(n1)  
	fmt.Println("n1=",n1)

	dec,_ := test2(12, 11)
	// fmt.Println("sum=",sum)
	fmt.Println("dec=",dec)

	// 指针在函数中传值
	var zzz int =10 
	// Wer(&zzz)
	fmt.Println("原来的zzz=",zzz)

	// 在go中，函数也是一种数据类型，
	// 可以赋值给一个变量，通过变量名来直接调用函数
	fun := test  // 变量fun 接收 test函数 
	fun(4)  // 调用

	// 函数可以作为参数，传入函数中
	var res int = myFun(test, 10)
	fmt.Println("__-res=",res)

	


	
}