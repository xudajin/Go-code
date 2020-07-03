package main
import "fmt"
// 定义域全局变量
var d1 = 300
var d2 = "luffy"
// 定义全局变量的方式2 
var(
	q1 = 400
	q2 = "xx"
)

func main()  {
	//定义变量方式一
	var i int = 10  // var 关键字， i 变量名 int 变量类型
      				//  将变量i 赋值为10，若不赋值int类型的默认值为0
	// 使用变量
	fmt.Println("i=",i)   // 打印i的值

	// 定义变量 方式二 根据值自行判断变量类型
	var num =10    //类型推导
	fmt.Println("num=",num)

	// 定义变量 方式三 	省略var，使用 :=  符号来定义变量 
	name := "tom"
	fmt.Println("name=",name)	

	// 多变量声明
	var n1, n2 ,n3 int  //不对变量赋值，则int类型默认为0
	fmt.Println("n1=",n1,"n2=",n2,"n3=",n3)
	// 多变量声明2
	var c1,nick = 100,"jack"
	fmt.Println("c1=",c1,"nick=",nick)
	// 多变量声明3   类型推导
	f1 ,city := 200 ,"北京"
	fmt.Println("f1=",f1,"city=",city)

	// 输出全局变量
	fmt.Println("d1=",d1,"d2=",d2)
	fmt.Println("q1=",q1,"q2=",q2)
}