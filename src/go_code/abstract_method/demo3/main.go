package main
import (
	"fmt"
)
//继承的细节
type A struct {
	Name string
	age int
}
func (self *A) SayOK() {
	fmt.Println("A sayok",self.Name)
}
func (self *A) hello() {
	fmt.Println("A hello",self.Name)
}

// 定义B 继承A
type B struct {
	A
}


func main()  {
	// 注意：结构体可以使用嵌套匿名结构体所有方法和字段
	//不管首字母是大写还是小写，都可以使用
	var b B	
	b.A.Name ="railgun"
	b.A.age = 19   //可以访问
	b.A.SayOK()
	b.A.hello()  // 方法名小写，可以使用
	// 访问匿名结构体的写法可以简化 
	// b.A.age 等价于 b.age  编辑器会自动去寻找对应的值
	//使用b.age 若b中没有age属性会去继承的结构体中找age属性
	//属性和方法的访问遵照就近原则，当B和A 同时都有sex属性时
	//b.sex 会优先访问B中的属性，若想访问A中的sex属性可以使用
	//b.A.sex 进行区分
}