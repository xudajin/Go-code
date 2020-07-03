package main

import (
	"fmt"
)

// 遍历数组
func main()  {
	arr := [3]int{1,2,3}
	for i ,v := range arr {
		fmt.Println(i,v)
	}

	
}