package main

import (
	"fmt"
)

func main()  {
	// var second float64 = 5
	// var sex string ="woman"

	// if second < 8 {
	// 	if sex == "man" {
	// 		fmt.Println("进入男子组决赛")
	// 	} else {
	// 		fmt.Println("进入女子组决赛")
	// 	}
	// } else {
	// 	fmt.Println("无法进入决赛")
	// }
	var season int = 11
	var age int =15

	if season >= 4 && season <= 10 {
		if age < 18 {
			fmt.Println("儿童票30")
		} else if age > 60 {
			fmt.Println("老人票20")
		} else {
			fmt.Println("成人票60")
		}
	} else {
		if age > 18 && age <60 {
			fmt.Println("成人票40")
		} else {
			fmt.Println("其他20")
		}
	}


	
	


}





