package main

import (
	"fmt"
)


// 切片是用于针对对不确定长度的数据类型
// 切片是数组的引用 则 切片是引用类型
// 切片的长度是可以动态变化的
func main()  {
	var arr = [4]int{1, 2, 3 , 4} // 数组
	// 定义一个切片
	// slice := arr[1:3]  表示引用arr数组，引用下标1到3
	// 不包含下标3 的值：
	mySlice := arr[1:3]  //在数组上进行切片
	//定义切片方法二：
	var newSlice = []int{3,4,5,6}

	fmt.Println(arr)
	fmt.Println(mySlice) // 切片的容量可以变化
	fmt.Println(newSlice)

	// 若声明切片不进行赋值的话，零值为nil，
	// 且无法使用 slice[0]==xx 赋值，此时切片的长度为0
	// 故无法使用下标操作

	//切片的长度和容量
	s := []int{1,2,3,4,5,6}
	fmt.Printf("长度为=%d  容量为：%d\n",len(s),cap(s))

	a := s[2:3]  //对切片进行切片  长度1  容量4 （3，4，5，6） 
	fmt.Printf("长度为=%d  容量为：%d\n",len(a),cap(a))
	// 当一个切片是在一个原切片或数组上产生的时，
	// 长度和容量就会不同，容量为新切片的第一个元素，到其原
	//数组的末尾的个数
	
	//切片的本质就是对底层数组的封装
	// 它包含三个信息：底层数组的指针，切片的长度，和切片
	//的容量

	//使用make()函数创建一个切片
	var sliceA = make([]int, 4, 8) //(int,4,8) 类型，长度，容量
	fmt.Printf("%d----%d",len(sliceA),cap(sliceA)) // 4---8


}