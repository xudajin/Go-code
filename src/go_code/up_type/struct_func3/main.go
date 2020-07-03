package main

import (
	"fmt"
)
type Person struct {
	Name string
}
//给Person增加方法
func (self Person) speak() {
	fmt.Println(self.Name,"是个好人")
}
func (self Person) add(n1 int, n2 int) int {
	return n1 + n2
}
// 当传入的参数是 结构体的指针时 ,self 默认传入的是实例对象的地址
func (self *Person) add2(n1 int, n2 int) int {
	return n1 + n2
}
//方法调用时的参数传递是值拷贝
func main()  {
	var p Person  //声明变量
	p.Name = "上条"  //赋值
	p.speak()	//调用方法
	res := p.add(10,20)  // 调用方法
	fmt.Println(res)
	res2 := (&p).add2(50,20)  // 调用方法
	fmt.Println(res2)
}	





