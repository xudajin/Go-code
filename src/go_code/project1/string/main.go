package main

import (
	"fmt"

)

// 字符串就是一串固定长度的字符连接起来的字符序列
// go的字符串是由单个字符串联起来的

func main()  {
	//string 类型的使用
	var addr string = "北京长城" //字符串使用双引号
	fmt.Println(addr)

	// go中的字符串是不可变类型的
	// 字符串一旦赋值，字符串就不能修改
	var str string ="hello"
	//str[0] = "a" //报错 不能通过索引去修改字符串
	fmt.Println(str)
	//  字符串的两种表示形式 1：双引号，会识别转义字符 2 反引号，
	// 以字符串的原生形式输出，包括换行和特殊字符，可以实现防止攻击
	// 输出源代码等效果

	var str2 string = "abc\nacb" // 换行输出
	fmt.Println(str2)

	// 使用反引号``，将`` 内的所有字符识别为字符串

	var str3 string = ` 
	package main

	import (
		"fmt"
	
	)
	
	// 字符串就是一串固定长度的字符连接起来的字符序列
	// go的字符串是由单个字符串联起来的
	
	func main()  {
	`
	fmt.Println("str3",str3)
	// 字符串的拼接
	var str4 string ="hello" + "world"
	str4 +=" hehe"   // 拼接
	fmt.Println(str4)

	// go中的基本数据类型的默认值
	//go中的所有数据类型都有默认值叫零值，当没有对变量赋值时
	// 默认使用零值
	//  类型     默认值
	//  int       0 
	//  float     0 
	//  string    ""       空字符串
	//  bool      false








	
}