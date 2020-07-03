package main

import (
	"fmt"
)

func main()  {
	// 创建map类型数据
	// 类似于python字典 ////

	var myMap =make(map[string]string)

	myMap["zzz"] ="fdfd"
	myMap["age"]="23"
	fmt.Println(myMap["age"])
	//方法二
	var map1 = map[string]string{
		"name":"路飞",
		"age":"20",
	}
	fmt.Println(map1)

	for k,v := range map1 {
		fmt.Println(k,v)
	}

	//如何判断map中的某个键是否存在  
	v,exists := map1["nfdf"]  //当键不存在时 exists =false
	fmt.Println(v,exists)  

	// 删除map中的key以及value
	delete(map1,"age")  //删除"age
	fmt.Println(map1)  

}