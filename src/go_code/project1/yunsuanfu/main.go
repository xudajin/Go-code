package main



import (
	"fmt"
)
// 算术运算符


func main()  {
	//重点讲解 /  % 
	// 说明，如果运算的数都是整数，那么除后，去掉小数部分
	// 保留整数部分
	
	fmt.Println(10 / 4) // 输出2 

	var n1 float32 = 10 / 4  //输出 2 
	fmt.Println(n1)

	// 如果希望保留小数部分，则需要浮点数参与运算
	var n2 float32 = 10.0 / 4   //输出2.5
	fmt.Println(n2)
	

	// %（取余） 的使用
	// 公式 ：  a % b = a -a/ b * b 


	var i int =10 
	i ++    // i =i + 1  ，i ++
	// i -- 同 i ++
	fmt.Println(i)	//输出 11

	// i ++ 只能独立使用 
	// var n int 
	// n = i ++   // 报错 i++ 独占一行，无法在添加其他代码

	// 练习题:
	// 还有97天放假 ，问：还有多少星期，多少零天

	var holiday int =97
	var week int = holiday / 7 
	var day int = holiday % 7 
	
	fmt.Printf("还有%d星期零%d天",week,day)

	fmt.Println("关系运算符————————————")


	// 关系运算符（比较运算符）
	//  ==  , !=  < , > , <= ,>= 
	var c1 int =9 
	var c2 int =8 
	fmt.Println(c1 == c2)
	fmt.Println(c1 != c2)
	fmt.Println(c1 > c2)
	fmt.Println(c1 < c2)
	fmt.Println(c1 >= c2)
	fmt.Println(c1 <= c2)
	var flag bool = c1 >c2
	fmt.Println("flag",flag)

	// 逻辑运算符
	// &&  与 
	// ||  或 
	// !  非  !(ture|false)

	var age int = 40 
	if age > 30 && age< 50 {
		fmt.Println("ok")
	}
	if age > 30 && age <=40 {
		fmt.Println("ok2")
	}
 



	
	


}






