package main
import "fmt"
//嵌套匿名结构体后，可以直接在创建结构体变量时
//直接指定匿名结构体字段的值

type School struct {
	Name string
}
type Skill struct {
	Name string 
	Hit  int
}

type User struct {
	Name string
	School
	Skill 
}

func main()  {
	//嵌套匿名结构体后，可以直接在创建结构体变量时
	//直接指定匿名结构体字段的值
	user := User{
		"misaka",
		School{"常盘太"},
		Skill{"某科学的超大头炮",100},
	}
	fmt.Println("user",user)




}