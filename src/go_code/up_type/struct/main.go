package main

import "fmt"

//golang中没有 "类"的概念，Golang中的结构体和
//其他语言的 类 有点相似
//且 GO中的结构体具有更高的扩展性和灵活性


// 定义结构体的格式
/*
type 类型名 struct {     //类型名 为结构体的名称，在同一个包内部能重复
	字段名 字段类型      // 字段名结构体字段的名字，字段名是唯一的 
	字段名 字段类型      // 字段类型 结构体字段的具有类型
	字段名 字段类型
}
*/ 
// 构造一个Person的结构体：方法一
// 首字母小写表示私有，首字母大写表示公有
// ps : 在go中支持对 结构体指针 直接使用 . 来访问
//结构体成员。p1.name ="张三" 其实在底层是(*p2).name ="张三"
type Person struct {
	name string
	age int
	sex string
}
// 结构体是值类型的数据 修改一个实例化对象，不会影响另一个实例化对象
func main() {
	// 要使用结构体，需要进行实例化
	// 方法一：
	var p1 Person //实例化结构体
	p1.name = "死神"
	p1.age = 30
	p1.sex = "女"
	fmt.Printf("值:%#v 类型:%T\n",p1,p1) 
	//打印结果 值:{死神 30 女} 类型:main.Person

	//方法二：使用new(结构体名) 来实例化一个对象
	var p2 = new(Person)
	p2.name = "哈哈"
	//使用指针来修改结构体中的成员
	(*p2).age = 40
	fmt.Printf("值：%#v  类型%T\n", p2,p2)
	//打印结果 值：&main.Person{name:"哈哈", age:40, sex:"", work:""}  类型*main.Person

	//方法三 ：实例化结构体
	var p3 = &Person {}
	p3.name ="呵呵"
	p3.age = 32
	p3.sex = "女"
	fmt.Printf("值：%#v  类型%T\n", p3,p3)
	// 输出结果 值：&main.Person{name:"呵呵", age:32, sex:"女"}  类型*main.Person

	// 方法四
	var p4 = Person{
		name: "railgun",
		age: 20,
		sex : "男",
	}
	fmt.Printf("值：%#v  类型%T", p4,p4)
	// 结果 值：main.Person{name:"railgun", age:20, sex:"男"}  类型main.Person



}