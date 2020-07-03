package main

import (
	"fmt"
	"encoding/json"
)

// 在结构体的每个字段上 可以写个tag
// 该tag可以通过反射机制获取
// 常用与序列化和反序列化  如将struct 变量转成json
 
type Monster struct {   //`tag` 结构体标签
	Name string `json:"name"`
	Age int `json:"age"`
	Skill string `json:"skill"`
}

//当 结构体的属性是小写时，其他包的函数无法使用
func main()  {
	//实例化Monster
	monster := Monster{"悟空",600,"金箍棒"}
	// 将monster 序列化成一个json格式的字符串
	jsonstr,err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json 错误",err)
	}
	fmt.Println(string(jsonstr))



	
}