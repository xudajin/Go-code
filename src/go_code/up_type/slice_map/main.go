package main
import (
	"fmt"
)

// map是引用类型的数据 所以会影响原map的数据
func main()  {
	//创建一个切片，切片的元素是map
	var info = make([]map[string]string, 3, 3)
	fmt.Println(info[0])  //map[] map的默认值nil
	//声明info切片中的第一个元素为map类型的数据
	info[0] = make(map[string]string) 
	info[0]["name"] = "死神" // 对map赋值
	info[0]["age"] = "10"  
	fmt.Println(info)
	// 打印输出:[map[age:10 name:死神] map[] map[]]


	// 声明一个map ，里面嵌套切片
	var useinfo = make(map[string][]string)
	useinfo["like"] = []string{
		"吃饭",
		"睡觉",
		"敲代码",
	}
	useinfo["work"] = []string{
		"python",
		"mysql",
		"linux",
		"golang",
	}
	fmt.Println(useinfo)


}