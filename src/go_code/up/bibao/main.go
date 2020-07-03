package main

import (
	"fmt"
	"strings"
)

// 闭包： 一个函数和与其相关的引用环境组合的一个整体


// 累加器 （闭包）
func Add() func(int) int { //Add函数返回一个func
	var i int = 5  // 外部变量
	return func(x int) int {   // 内函数
		i = i+ x 
		return i
	}

// 返回的是一个匿名函数，但匿名函数引用到了函数外的i
// 因此这个匿名函数和变量i形成一个整体，构成闭包
}

func makeSuffix(suffix string) func(string) string {
	// 如果 fileName 没有后缀，则加上
	return func(fileName string) string {
		if !strings.Contains(fileName,suffix) {
			return fileName + suffix
		}
		return fileName
		
	}
}


func main()  {
	// 闭包中的 i 的值会保存在内存中，当再次调用函数时
	// 不会重新定义i ，而是直接 从内存中获取i的值
	// 定义闭包类似于pytohn 中定义类 ，i 为类属性，调用
	// 闭包函数则为 实例化对象调用方法 对类属性进行操作
	res := Add()
	fmt.Println(res(1))  // 6  i=6 保存在内存中
	fmt.Println(res(2))  // 8    i + 2  =8 
	fmt.Println(res(3))  // 11    8 + 3 = 11
	

	f2 := makeSuffix(".avi")
	fmt.Println("文件名称=", f2("railgun.avi"))

	// 练习





}