package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"  // 时间相关的函数
)

// 内建函数的使用
func main()  {
// 字符串的相关内建函数
	// len()  统计字符串的长度，按字节  汉字占3个字节
	var str string = "dfafasfa"
	var num int = len(str)
	fmt.Println(num)

	// 字符串转整数  n，err := strconv.Atoi("12")
	n, err := strconv.Atoi("1d23")
	if err != nil {    //nil 为零值
		fmt.Println("转换错误")
	} else {
		fmt.Println(n)
	}

	// 整数转字符串  str = strconv.Itoa(123456)
	var str1 string = strconv.Itoa(123456)
	fmt.Printf("str= %v, str = %T\n",str1,str1)

	// 查找子串是否在指定的字符串中 
	var Y bool = strings.Contains("ssssaa","qsa")
	fmt.Println(Y)


	//统计一个字符串中有几个指定的子串  返回个数
	var count int = strings.Count("ceheesa","e") 
	fmt.Println(count)

	// 不区分大小写的字符串比较(==是区分大小写的)
	var z bool = strings.EqualFold("abc","ABC") // true
	fmt.Println(z)

	// 替换  产生变换后的值，赋值给res  qw没有变化
	var qw string = "go is good"
	var res string  = strings.Replace(qw,"go","go语言",2)
	fmt.Println(res)

	// 拆分  []string 数组类型
	var ew []string = strings.Split("hello,world,ok",",")
	fmt.Println(ew)
	fmt.Printf("ew的类型%T",ew)  // []string

	fmt.Println("________________")
// 日期的相关函数    具体查询go的函数手册
	// 获取当前时间
	now := time.Now()
	fmt.Println(now)


	
// new的使用
	num0 :=new(int)   // 生成一个int的指针类型，用于分配内存
	fmt.Println(num0)  //0xc000010490



	
}