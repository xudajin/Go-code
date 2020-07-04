package main

import "fmt"
// 多继承
// 当 C 同时继承A,B 则在使用A,B都有sex字段时，需要
// 显示指定使用的是A的sex 还是B的sex

type A struct {
	Name string
	age int
}
type B struct {
	Name string
	age int
}

type C struct {
	A
	B
}

func main()  {
	//嵌入多个结构体
	var c C 
	c.A.Name = "tom"  //显示指定 A的Name 还是B的Name
	fmt.Println(c.A.Name)
}