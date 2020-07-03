package main
// 函数练习
import (
	"fmt"
)


// 请编写一个函数swap(n1 *int n2 *int) 可以交换n1 和 n2的值
func swap(n1 *int, n2 *int) (int,int) {
	// 定义临时变量
	var mid int
	mid = *n1  // *指针变量为值
	*n1 = *n2 
	*n2 = mid
	return *n1,*n2
	
}



func main()  {
	var num1 int =10
	var num2 int =20
	a ,b := swap(&num1,&num2)  //传入num1 和num2 的地址
	fmt.Println(a,b)






	
}