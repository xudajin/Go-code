package main

import "fmt"

type car struct {
	carName string 
	color string
	price float64
} 
//声明方法
func (self car) run(road int) {  // c 类似 self  car 结构体名
	fmt.Printf("%v汽车跑了%d公里\n",self.carName,road)
}
// 当想修改实例化对象中的属性值时，必须传入结构体指针
func (self *car) changeColor(newColor string) {
	self.color = newColor	
}
func main()  {
	var p1 = car{
		carName: "bmw",
		color : "yellow",
		price : 22.25,
	}
	p1.run(3)  //调用方法   bmw汽车跑了3公里
	p1.changeColor("pink")  //修改颜色
	fmt.Printf("修改后车的颜色是%v",p1.color)

// golang中的结构体实例是独立的，不会相互影响





}