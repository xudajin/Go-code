package main

import (
	"fmt"
)

// 数组和切片


//  传递 指针变量
func test01(arr *[3]int) {
	(*arr)[0] =22
	
}



func main()  {
// 定义数组
	// [n]int   n为数组中最多存放的元素数量
	// float64 元素的数据类型 ，若不赋值，使用零值
	var hens [6]float64  
	hens[0] = 1.2
	hens[1] = 5.2
	hens[2] = 2.2
	hens[3] = 4.2
	hens[4] =6.2
	hens[5] = 7.2
	fmt.Println("数组",hens)
	
// ps:数组的内存地址为数组中第一个元素的所在的地址
// &hen == &hen[0]
	fmt.Printf("%p\n",&hens)
	fmt.Println(&hens[0])
	
	// 初始化数组的四种方法
	var numArr01 [3]int = [3]int{1,2,3}
	fmt.Println("Arr1=",numArr01)

	var numArr02 = [3]int{3,4,5}
	fmt.Println("Arr2=",numArr02)
	
	var numArr03 = [...]int{7,8,9}
	fmt.Println("Arr3=",numArr03)
	
	var numArr04 = [...]int{1:100, 2:33, 3:66} //指定下标 1，2，3
	fmt.Println("Arr4=",numArr04)

	
// 数组的注意事项
	// 1 数组是多个相同类型的数据集合，一个数组一旦声明
	// 数组的长度是固定的，不能动态变化，否则报越界
	// 长度是数组的一部分，在传递函数参数时，需要考虑数组的
	//长度，长度不同会报错
	// 数组属于值类型，默认情况下是值出传递，因此进行值拷贝
	// 数组间不相互影像
	arr := [3]int{1,2,3}
	test01(&arr)
	fmt.Println(arr)
	
	// 定义多维数组
	var list = [3][2]string{
		{"北京","1"},      //每个数组不能在同一行
		{"上海","2"},
		{"成都","3"},
	}
	fmt.Println(list[2][1])
	for _,v := range list {
		for _,v2 :=range v {
			fmt.Println(v2)
		}
	}

}