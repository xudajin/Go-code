package main


import (

	"fmt"
)

// goto 语句可以无条件地转移到程序中指定的行
// goto 语句通常与条件语句配合使用，可用来实现条件转移
// 一般不使用goto语句， 以免造成程序流程的混乱

//基本语法
// goto label

func main()  {
	// 演示goto的使用,一般和if语句一起使用
	// 被跳过的代码在编辑器中会显示警告,但编译正常
	var n int = 30
	if n > 20 {  // 当满足时，直接执行label1 后的代码
		goto label1
	} 
	fmt.Println("1111")
	fmt.Println("2222")  
	label1:
	fmt.Println("3333")
	fmt.Println("4444")

	
}