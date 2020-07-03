package main

import (
	"fmt"
)
// swith 分支  匹配项中不需要加break，case后默认break
func main()  {
	// case 后的表达式可以有多个，用，隔开
	var i byte = 'f'
	switch i {
	case 'a':
		fmt.Println("周一")
	case 'b':
		fmt.Println("周二")
	case 'c':
		fmt.Println("周三")
	default:
		fmt.Println("输入有误")
		
	}





}