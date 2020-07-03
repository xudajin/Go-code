package main

import (
	"fmt"
	"errors"
)


func test()  {
	//使用defer + recover 来处理异常
	defer func()  {   // 定义一个匿名函数 来捕获异常
		err := recover()  // recover()捕获异常
		if err !=nil {   //err 不为nil ，则存在err
			fmt.Println("出现错误：",err)
		}
	} ()  // 直接调用 捕获异常的匿名函数
	var num1 int =10 
	var num2 int = 0
	var res int = num1/num2     // 10/0报错
	fmt.Println("res的值=",res)  //该行代码不执行
	
}

//案例
func readFile(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("读取配置错误")
	}
	
}

func test2()  {
	err := readFile("config.ni")
	if err != nil {
		panic(err)   //panic 抛出自定义错误，并停止程序
	}
	fmt.Println("test02 执行")
	
}


// 定义错误

func chu(a int, b int) int {
	defer func ()  {
		err := recover()
		if err != nil {
			fmt.Println("发送邮件")	
		}
	} ()
	return a/b
}


func main()  {
// go中的错误处理
// 默认情况下，当发生错误后(panic)，程序就会退出
// 所以，当发生错误后，可以捕获错误，并进行处理，保证
// 程序正常运行
// 错误处理的处理方式  defer， panic ， recover
// go可以抛出一个panic的异常，然后再defer中通过recover
//捕获这个异常，然后正常处理

	// test()
	// fmt.Println("错误处理完成")

// 错误处理后自定义错误信息
// errors.New("错误说明")   返回一个error类型的值，表示错误
// panic 内置函数,接收一个interface{} 类型的值 作为参数

//案例：
	// test2()
	chu(1,0)
	fmt.Println("aaaa")


// recover 只在defer 调用函数时才有效







	
}