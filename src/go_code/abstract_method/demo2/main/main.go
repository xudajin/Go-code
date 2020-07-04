package main

import (
	"fmt"
)
type Student struct {
	Name string
	Age int 
	Score int
}
func (self *Student) ShowInfo() {
	fmt.Printf("名字：%v  年龄：%v   分数:%v\n",self.Name,self.Age,self.Score)	
}
func (self *Student) SetScore(score int) {
	self.Score = score
}
// 继承 Student
type Pupil struct {
	Student   // 嵌入Student 结构体 既继承
}

func (self *Pupil) Testing() {
	fmt.Println("小学生正在考试")
}
// 继承 Student
type Railgun struct {
	Student   // 嵌入Student 结构体 既继承

}
func (self *Railgun) Testing() {
	fmt.Println("大学生正在考试")
}
func main()  {
	// 面向对象的继承
	// 当我们嵌入了匿名结构体后 使用方法
	pupil := &Pupil{}
	pupil.Student.Name =" tom"
	pupil.Student.Age = 8
	pupil.Testing()
	pupil.SetScore(50)
	pupil.Student.ShowInfo()

	railgun := &Railgun{}  // 实例化对象，指向地址
	railgun.Student.Name =" railgun"
	railgun.Student.Age = 20
	railgun.Testing()
	railgun.SetScore(67)
	railgun.Student.ShowInfo()












}