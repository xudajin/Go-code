package utils
import (
	"fmt"
)

// 为了让其他包使用，必须将函数名首字母大写
func Cal(n1 float64, n2 float64, do byte) (float64) {
	var res float64
	switch do {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2 
	case '*': 
		res = n1 * n2

	default:
		fmt.Println("输入有误")
	}
	return res
}