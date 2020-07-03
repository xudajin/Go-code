package main
import "fmt"

func test(n1 int) int {
	fmt.Println("test n1=",n1)
	return n1 +2
}
// 下面代码表示: myFun 等价于 func(int, int) int
type myFun func(int) int 

func newFun(fun myFun,num1 int) int {
	return fun(num1)
}

// go中支持对函数的返回值命名       //对返回值命名
func getSumAndSub(n1 int ,n2 int) (sum int , sub int) {
	sum = n1 + n2
	sub = n1 - n2 
	return sum ,sub
}

// GO支持可变参数  // args... 为接收所有int 参数
// 可变参数的格式为 变量名...  一般使用args...
func test4(args... int) int {  // 可变参数必须在最后
	var sum int 
	// 遍历args  
	for _ ,v := range args {
		sum +=v 
	} 
	return sum	
}

// 为了简化数据类型的定义，Go支持自定义数据类型

// 基本语法  type 自定义类型名 数据类型 // 相当于一个别名
func main()  {
	// 案例
	type myInt int // 定义一个类型
	var num1 myInt  //声明一个 myInt 类型的变量
	var num2 int
	num1 = 40
	num2 =int(num1)  //将myInt 类型 强转成int类型
	fmt.Println("num1=",num1, "num2=",num2)

	//案例二
	
	var res int =newFun(test,10)
	fmt.Println("__--res=",res)


	// go中支持对函数的返回值命名
	a,b := getSumAndSub(10,20)
	fmt.Println(a,b)

	
	var c int =test4(4,2,1,6)
	fmt.Println("c=",c)

}