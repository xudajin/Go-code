package main
import "fmt"

func main()  {
	//输出10句 hello
	// for 循环
	// 方法一 
	for i := 1; i <= 10; i++ {
		fmt.Println("hello")
	}

	//方法二  
	var j int =1   //变量初始化
	for j <= 10 { // 条件语句
		fmt.Println("world!")
		j++  // 循环变量迭代
	}

	// 死循环,一般配合break使用
	var k int = 1 
	for {  // 没有判断语句，死循环
		if k <= 10 {
			fmt.Println("ok")
		} else {
			break //跳出循环
		}
		k++  // 循环变量添加
	}
	
	// 对字符串进行遍历
	//方式一：
	var str string = "hahahaha"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n",str[i])
	
	}
	// 如果字符串含有中文，这方式一无法使用，方式一是按照
	// 字节来遍历的，在utf-8编码中，一个汉字占3个字节，
	//所以会产生乱码
	// 解决方法 ： 将str 转成 []rune 切片 


	
	//方式二： 不会存在中文遍历报错的问题
	for _ ,val := range str {  // for - range
		fmt.Printf( "val=%c\n",val)
	}

	// 练习题
	var count int = 0 
	var total int = 0
	for i:= 1; i <=100 ; i++ {
		if i % 9 == 0 { 
			fmt.Println(i)
			count++
			total += i
		}
	}
	fmt.Printf("count=%d total=%d",count,total)

	var num int =60
	for i := 0; i<= num ; i++ {
		fmt.Printf("%d + %d = %d\n",i,num-i,num)
	} 

	

}