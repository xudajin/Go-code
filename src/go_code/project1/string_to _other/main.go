package main

import (
	"fmt"
	"strconv"
)

// string转成其他基本数据类型

func main()  {
	var str string = "true"
	var b bool
	// 1, strconv.ParseBool 会返回两个值，所以用b接收
	// 第一个值，用 _ 来忽略掉第二个值  
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T , b=%v\n",b,b)

	var str2 string = "1234560"
	var n1 int
	// n1, _ =strconv.ParseInt(str2,10,64)
	n1, _ = strconv.Atoi(str2)
	fmt.Printf("n1 type %T , n1=%v\n",n1,n1)

	// 当 将 string "hello" 转成 int 时，出现错误，
	// 会转化值的类型，但值为0 
	//当类型转换失败时，类型改变，值会重置成该数据类型的零值


	
}
