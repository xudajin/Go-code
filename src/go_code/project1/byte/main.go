package main

import (
	"fmt"
)

//go中没有专门的的字符类型，如果要存储单个字符(字母)，
//一般使用byte来保存

// 字符串就是一串固定长度的字符连接起来的字符序列。
//go的字符串是由单个字节连接起来的，也就是说对于传统的
//字符串是由字符组层的，而go的字符串不同，它是由字节组成

func main()  {
	// 当我们直接输出byte值，就是输出了其对应的字符的ascll码
	//的码值
	// byte 的范围在0~255之间
	var c1 byte = 'a'  //97
	var c2 byte = '0'  //48
	fmt.Println(c1, c2)
	// 如果我们希望输出对应的字符，需要格式化输出
	fmt.Printf("c1=%c c2=%c\n",c1 ,c2)

	//var c3 byte = '北' //报错:北的码值21271 超出byte的码值范围
	// 在go中 单引号代表byte类型， 双引号才表示字符串
	var c4 int = '北' // int的范围大于21271 索引可以保持'北'
	fmt.Printf("c4=%c ， %T",c4,c4) // 控制台输出 北

// 字符常量是用单引号('')。例如 var c1 byte = 'a'
// 在go中，字符的本质是整数，直接输出时显示该字符对应的
//ascll 的编码值
// 在utf-8编码中，英文字母是1个字节，汉字是3个字节
// 字符类型对应ascll表的码值，相当于一个整数，所以可以
//进行加法运算，即为字符的ascll码值相加
// 字符类型 存储和读取的过程
// 存储 ：字符 -> 对应码值 ->二进制 ->存储
// 读取: 二进制 -> 码值 ->字符 -> 读取




	
}
