package main

import (
	"fmt"
	"strconv"
)

// 基本数据类型与string类型的转换
func main()  {
	
	var num1 int =99
	var num2 float64 = 23.456
	var b bool = true
	var myChar byte = 'h'  //byte 类型

	var str string  // 未赋值，空字符串

	// 方式一：使用fmt.Srintf 方法转换
	str = fmt.Sprintf("%d",num1)
	fmt.Printf("type %T , %v\n",str,str)

	str = fmt.Sprintf("%f",num2)
	fmt.Printf("type %T , %v\n",str,str)

	str = fmt.Sprintf("%t",b)
	fmt.Printf("type %T , %v\n",str,str)

	str = fmt.Sprintf("%c",myChar)
	fmt.Printf("type %T , %v\n",str,str)
	fmt.Println("---------方法二-------")
	// 方式二：使用strconv 包的函数

	var w1 int = 99 
	var w2 float64 = 23.56
	var w3 bool = true

	// int64(w1) 为将w1强转成int64类型 ， 10 表示10进制
	str = strconv.FormatInt(int64(w1),10)
	fmt.Printf("type %T ,  str = %q\n",str,str)
	// strconv 包中的Itoa函数
	var e1 int = 4531
	str = strconv.Itoa(e1)
	fmt.Printf("e1 type %T ,str=%q\n",str,str)


	//  w2 为参数, 'f' 为格式  10为 小数位保留10位 64表示float64
	str = strconv.FormatFloat(w2,'f',10,64)
	fmt.Printf("type %T  , str= %q\n",str ,str)
	// bool 转 str
	str = strconv.FormatBool(w3)
	fmt.Printf("type %T ,str=%q",str,str)

	
}