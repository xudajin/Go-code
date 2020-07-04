package main

import "fmt"

//如歌一个结构体嵌套了一个有名结构体，这中模式叫组合
// 如果是组合关系，那么在访问组合的结构体的字段或方法时
//必须带上结构体的名字
type D struct {
	Name string
}

type F struct {
	D  
}
type G struct {
	d D  // 对结构体D 命名为d 
}
func main()  {
	var f F  // 开辟f 空间
	var d D   // 开辟d空间   
	f.Name ="FFF"
	fmt.Println(f.Name)   //输出 FFF
	fmt.Println(f.D.Name)   //输出 FFF
	fmt.Println("--",d.Name)   //输出 nil
	//我们声明对象f时 f.Name 和 f.D.Name 是在同一个空间中
	// 使用f.Name="FFF"时 ，等价于对f.D.Name 赋值
	//而当我们在再次声明对象d时，就开辟了新的空间
	//而在d对象的空间中 我们没有对d.Name 赋值，故输出 nil
	//在f空间中对D的操作 不会影响到d空间的属性
	//ps  因为结构体实例是独立的 故 对象d 不等于 f 的父对象
	var g G
	// fmt.Println(g.Name) // 显示undefine Name
	fmt.Println("--",g.d.Name) //g.别名.属性

}