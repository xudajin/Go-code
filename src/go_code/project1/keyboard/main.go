package main


import (
	"fmt"
)



//需要获取用户输入的数据，就可以使用键盘来获取
//键盘输入语句

func main() {
	// 要求：可以从控制台接收用户输入信息 name，age，salary
	// 是否通过考试

	// 方式1 fmt.Scanln
	var name string
	var age int8 
	var sal float32
	var isPass bool 
	fmt.Println("请输入姓名")
	// 当程序执行到 fmt.Scanln时 会阻塞 类似 input
	fmt.Scanln(&name) // 传递name的内存地址

	fmt.Println("请输入年龄")
	fmt.Scanln(&age)
	
	fmt.Println("请输入薪水")
	fmt.Scanln(&sal)

	fmt.Println("请输入是否通过考试")
	fmt.Scanln(&isPass)

	fmt.Printf("名字是%v  年龄是%v 薪水是%v  是否通过考试%v ",name,age,sal,isPass)


	//方法二 fmt.Scanf

	fmt.Printf("请输入你的姓名 年龄 薪水 是否通过考试，使用空格隔开")
	fmt.Scanf("%s %d %f %t",&name, &age ,& sal, &isPass)
	fmt.Printf("名字是%v  年龄是%v 薪水是%v  是否通过考试%v ",name,age,sal,isPass)

}
