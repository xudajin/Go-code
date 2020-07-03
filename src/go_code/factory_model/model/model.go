package model

//定义一个结构体

type student struct{
	name string
	score float64
}
// 因为student结构体首字母是小写，因此只能在model中使用
// 我们通过工厂模式来解决

func NewStudent(Name string, Score float64) *student {
	return &student{
		name : Name,
		score : Score,
	}
}
func (self *student) GetScore() float64 {
	return self.score
}