package main

//go中的包的引用和管理
// 包的本质就是创建不同的文件夹，来存放程序文件

//包的基本概念 ：go的每一个文件都属于一个包，也就是说gi是以
// 包的形式来管理文件和项目目录结构的


//包的三大作用
/*
1 区分相同名字的函数，变量和标识符(不同包下,可以有相同的名称的变量或函数）
2 当程序文件很多时，可以很好的管理项目
3 控制函数，变量访问范围，既作用域
*/

// 打包 package 包名
// 引入包  import "包名路径"
// 引用 包名.函数名()
// 引用时不同包的函数时，需要函数名手字母大写，才可以引用


// 导入utils包中的函数
import (
	"fmt"
	"go_code/project1/package_info/utils" // 引入utils包中的函数
)
func main()  {
	// 案例
	// 导入 package utils 中 的Cal函数 
	var res float64 = utils.Cal(1.2 ,2.2 ,'+') 
	fmt.Printf("res=%.2f ",res)   // .2f 保留两位小数


}