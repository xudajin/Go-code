package main

import (
	"fmt"
)
// 结构体方法的练习
type MethodUtils struct {
	Name string
}

func (self MethodUtils) Method(m int, n int) {
	for i := 0; i<m; i++ {
		for j := 0; j<n; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
func (self MethodUtils) Area(m int, n int) int {
	return m * n
}

func (self *MethodUtils) Num(m int) {
	if m % 2 !=0 {
		fmt.Println(m,"是基数")
	} else {
		fmt.Println(m,"是偶数")
	}	
}


func main()  {
	var m =  MethodUtils {
		Name : "练习",
	}
	m.Method(5,10)
	res := m.Area(10,6)
	fmt.Println(res)
	(&m).Num(5)
	(&m).Name = "结构体练习"
	fmt.Println(m.Name)


	
}