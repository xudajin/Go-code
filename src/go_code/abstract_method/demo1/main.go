package main


import (
	"fmt"
)
//定义一个结构体

type Account struct {
	AccountNo string
	Pwd string
	Balance float64
}

//方法
func (self *Account) Deposite(money float64, pwd string) {
	// 输入的密码是否正确
	if pwd != self.Pwd {
		fmt.Println("你输入的密码错误")
		return 
	}
	// 输入的金额是否正确
	if money <= 0 {
		fmt.Println("你输入的余额错误")
		return 
	}
	self.Balance += money
	fmt.Println("存款成功")
}

func (self *Account) WithDraw(money float64, pwd string) {
	// 输入的密码是否正确
	if pwd != self.Pwd {
		fmt.Println("你输入的密码错误")
		return 
	}
	// 输入的金额是否正确
	if money <= 0 || money > self.Balance {
		fmt.Println("你输入的余额错误")
		return 
	}
	self.Balance -= money
	fmt.Println("取款成功")
}

func (self *Account) Query(pwd string) {
	if pwd != self.Pwd {
		fmt.Println("密码错误")
		return 
	}
	fmt.Printf("你账号的余额为%v\n",self.Balance)
}

func main()  {
	var account = Account{
		AccountNo : "gs123456",
		Pwd : "123456",
		Balance : 100.0,
	}
	
	account.Deposite(100.0,"123456")
	account.Query("123456")
	account.WithDraw(50.0,"123456")
	account.Query("123456")


}