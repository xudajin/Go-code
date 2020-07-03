package main

import (
	"fmt"
)

// 基本数据类型的相互转换
// go中的数据类型转换需要显示转换，不能自动转换

// 表达式T(v) 将值 v 转换为类型 T 

func main()  {
	var i int32 = 100 
	//希望将 i => float
	var n1 float32 = float32(i)  //float()强转函数
	var n2 int8 = int8(i)
	// %v 以原值输出
	fmt.Printf("i=%v n1=%v n2=%v\n",i ,n1,n2)

	// go中的数据类型转换 可以从范围小->范围大，范围大->范围小
	// 当 使用 var n3 float32= float32(i) 强转i时，生成一个
	// 一个新的值并赋值给n3 ，i 本身的数值和类型并没有发生改变
	
	// 在转换中，比如将 int64 转换成 int8 ，编译时不会报错
	// 只是转换的结果是按溢出处理，结果和我们想的不一样
	var num1 int64 = 999999
	var num2 int8 =int8(num1)  // 999999超出了int8的范围
	fmt.Println("num2=",num2)  //输出值为63，不是999999

	// 练习题 

	// var f1 int32 =12
	// var f2 int64
	// var f3 int8
	
	// f1 为int32 加上一个具体数值后 还是 int32类型
	// f2 是 int64类型 ,故 无法将int32的值赋值给int64
	//f2 = f1 +20   报错 
	// f3 =f1 +20   f3 与f2 同理 ，类型不匹配 int8 不匹配int32


	// 练习题2

	var d1 int32 = -12
	// var d3 int8
	var d4 int8 
	// 12+ 127 超出int8的范围，所以编译通过，但数值溢出
	//d3 = int8(d1)+127  
	// int8最大值为127 所以128 已经超出范围，编译无法通过
	d4 = int8(d1) +127
	fmt.Println(d4)




	
}