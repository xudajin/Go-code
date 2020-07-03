package main

import (
	"fmt"
	"go_code/factory_model/model"
)
// go的结构体中没有构造函数，通常使用工厂模式来处理
//当定义结构体时 结构体名的首字母是小写时，又希望其他包
// 能引用该结构体来创建实例， 这时就需要使用工厂模式



func main()  {
	// 创建model包下的Student结构体的实例
	// var stu = model.student{
		//  Name : "railgun",
		// Score : 99.9,
	// }
	//当model包中的结构体是首字母是小写时,编写函数NewStudent
	//返回结构体student的指针，并在main中调用函数，并获取
	// 函数返回的 结构体student实例的指针
	var stu = model.NewStudent("railgun",98.9)
	fmt.Println(stu,stu.GetScore())  //&{railgun 99.9}
}