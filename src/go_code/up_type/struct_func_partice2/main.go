package main



import (
	"fmt"

)
type Com struct {

}
func (self Com) add(m int, n int) int {
	return m + n
}
func (self Com) dec(m int, n int) int {
	return m - n
}
func (self Com) cheng(m int, n int) int {
	return m * n 
}
func (self Com) chu(m float64, n float64) float64 {
	return m / n 
}
func (self Com) all(m float64, n float64, do byte) float64 {
	res := 0.0
	switch do {
	case '+':
		res = m + n
	case '-':
		res = m -n	
	}
	return res
}

type Student struct {
	Name string 
	Gender string 
	Age int 
	Id int 
	Score float64
}
func (self *Student) say() string {
	info := fmt.Sprintf("学生信息:%v  %v  %v  %v  %v" ,
	self.Name,self.Gender,self.Age,self.Id,self.Score)
	return info
}

func main()  {
	var c Com
	fmt.Println(c.add(1,5))
	fmt.Println(c.dec(1,5))
	fmt.Println(c.cheng(1,5))
	fmt.Println(c.chu(1,5))

	res := c.all(4,10,'+')
	fmt.Println(res)
	var stu =Student{
		Name : "rail",
		Gender : "female",
		Age : 18,
		Id : 10,
		Score : 99.9,
	}
	fmt.Println(stu.say())
}