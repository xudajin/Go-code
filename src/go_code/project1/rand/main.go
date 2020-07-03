package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main()  {
	// 为了生成一个随机数，还需要给rand设置一个种子
	// time.Now().Unix() 返回当前时间的时间戳
	

	// break 和随机数的使用
	var count int = 0
	rand.Seed(time.Now().Unix())  // 随机种子放在循环外面
	for {
		var n int = rand.Intn(100)+1 // 生成随机数
		count++
		if n == 99 {
			break
		}
		
	}
	fmt.Println("一共用了",count)


	//break 语句出现在多层嵌套的语句中，可以通过标签来指定
	// 想要跳出的循环，如下
	lable1:  //设置标签
	for i := 0; i < 4 ; i++ {
		// lable2:
		for j := 0 ; j < 10 ; j++ {
			if j ==2 {
				break lable1  // 指定想跳出的标签
			}
			fmt.Println(j)
		}
	}

	fmt.Println("___________________")
	// 课堂练习
	var sum int 
	for i := 1; i <=100; i++ {
		sum +=i
		if sum >20 {
			fmt.Println("当前i的值为",i)
			break
		}
	}

	








}