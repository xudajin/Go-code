package main

import "fmt"


//递归的特点是 先进后出 ，最后的调用的函数最先打印出结果
func test(n int)  {
	fmt.Println("n=",n)
	if n > 2 {
		n--
		test(n)  // 函数点用函数本身		
	}		
}
// 斐波那契
func fibi(n int) (int)  {
	if n ==1 || n==2 {
		return 1
	} else {
		return fibi(n-1) + fibi(n-2) 
	} 
}

func part(n int) (int) {
	if n == 1 {
		return 3
	} else {
		return 2 * part(n-1) + 1
	}
}

func patch(day int) (int) {
	if day == 10 {
		return 1
	} else {
		return (patch(day+1) + 1) *2
	}
	


}


func main()  {
	// 函数递归
	test(5)
	// 斐波那契函数
	res := fibi(6)
	fmt.Println("res=",res)
	// 练习二   f(1)=3  f(n)= 2*f(n-1)+1
	num := part(5)
	fmt.Println("num=",num)
	// 桃子问题

	var patch int = patch(5)
	fmt.Println("桃子的个数",patch) 





	




	
}

