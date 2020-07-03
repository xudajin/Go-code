package main

import (
	"fmt"
	"errors"
)


// 循环打印输入的月份天数 (使用continue)
func date(year int, month int , day int) {
	if month <=0 || month >12 || day >31 {
		error := errors.New("参数输入错误") // 自定义错误
		panic(error)  // 抛出错误
	} else {
		fmt.Printf("日期是:%d-%d-%d",year,month,day)

	}
	 		
}

func susu() {
	var flag int =0
	var sum int =0
	for i := 1; i<100; i++ {
		if i % i == 0 && i % 1== 0 && i%2 !=0 && i%3 !=0 && i%5!=0 && i%7 != 0{
			fmt.Print(i," ")
			sum += i
			flag ++
		}
		
		if flag == 5 {
			fmt.Print(" sum=",sum)
			fmt.Println()
			flag = 0
			sum =0
		} 
	}
	

		
	
}

func main()  {
	// 练习一
	// date(2020,12,31)


	// 练习二
	susu()
	
}
