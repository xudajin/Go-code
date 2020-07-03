package utils


var Age int
var Name string

//Age 和Name 全局变量，我们需要在main.go 中使用

// init 函数完成初始化工作
func init() {
	Age = 100
	Name = "railgun"
}