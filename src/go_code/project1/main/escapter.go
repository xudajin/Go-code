package main // 属于main包，与是否同属于一个文件无关

import "fmt"  //输入输出的函数

func main()  {

	// fmt.Println("tom\tjk") // \t制表符

	// fmt.Println("tom \n jk") // \n 换行符

	// fmt.Println("tom \\ jk") // \\ 输出 \

	// fmt.Println("tomzz\rzz") // \ 当前行最前面开始输出，覆盖以前的内容
	fmt.Println("姓名\t年龄\t籍贯\t住址\njolin\t12\t河北\t北京")

	var num =2 
	var total = num +4 
	fmt.Println(total)
	
	
}

