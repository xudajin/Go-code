package main

import "fmt"

type Box struct{
	longer float64
	weight float64
	hight float64
}
func (self *Box) block() float64 {
	// 该self为指针，使用 (*self)获取指针的值
	// 也可以直接使用 self 编辑器底层进行优化
	block := (*self).longer * (*self).weight * (*self). hight
	return block
}

type Visitor struct {
	name string
	age int
}
// 接收的类型为指针  *Visitor
func (self *Visitor) price() int {
	var price int 
	if self.age > 18 {
		price = 20
	} else {
		price = 0
	}
	return price
}




func main()  {
	var box = Box{
		longer : 3.0,
		weight : 3.0,
		hight : 2.0,
	}
	res := box.block()
	fmt.Printf("%.1f\n",res)
	fmt.Println("---------------")
	

	var v = Visitor{
		name : "美琴",
		age : 9,
	}
	fmt.Println(v.price())
	// 实例化结构体的指针
	var v2 = &Visitor{"上条",20}
	fmt.Println(v2)  //  &{上条 20} 说明是指针

	

}

