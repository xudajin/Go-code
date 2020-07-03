package main

import "fmt"

//go中没有类的概念，但可以给类型定义方法。
//所谓方法就是定义了接收者函数
//接收者类似于python中的 self


//结构体方法的定义格式
/* 
func (接收者变量  接收者类型) 方法名(参数) 返回值 {
	
}
*/
type car struct {
	carName string 
	color string
	price float64
} 
//ps  结构体内部能存在函数，所以函数需要在结构体外定义
//声明方法 run方法与结构体car绑定,
// 所以只能通过 car结构体实例化的对象c 来调用
func (c car) run(road int) {  // c 类似 self  car 结构体名
	fmt.Printf("%v汽车跑了%d公里",c.carName,road)
}

func main()  {
	var p1 = car{
		carName: "bmw",
		color : "yellow",
		price : 22.25,
	}
	p1.run(3)  //调用方法   bmw汽车跑了3公里
	




}