package main


import (
	"fmt"
)
// 切片的修改
func main()  {
	var sliceA = make([]int ,4 ,6)
	sliceA[0] = 10 
	sliceA[1] = 20
	fmt.Println(sliceA)


	//golang中无法通过下标的方式对切片进行扩容
	// 对切片进行扩容需要使用append方法

	var sliceB []int //定义切片
	sliceB = append(sliceB,12)
	// 一次增加多个数值,切片发生扩容时，是在原基础上进行加倍扩容
	// 所长度和容量可能不一致，当发生扩容时，切片就不在指向
	// 原数组，而是开辟一个新空间
	// 切片是引用数据类型
	sliceB = append(sliceB,23,45,67,98)
	fmt.Printf("%v---%v --%v",sliceB,len(sliceB),cap(sliceB))


	// copy()方法，可以复制一个生成一个新的切片，修改原切片
	// 不在影响新的切片

	
}