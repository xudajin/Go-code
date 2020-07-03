package main


import (
	"fmt"
	"unsafe"

)

// 布尔类型中，值允许取值true 和false 
// bool 类型占1个字节

func main()  {
	var b bool =false
	fmt.Println(b)
	fmt.Println("b占用的空间=",unsafe.Sizeof(b))
	// 不可以使用0和非0的数值来 代替bool类型
	
}