package main

import "fmt"



func main()  {
	for i := 1; i <= 9; i++ {

		for k:=1; k<= 9-i; k++ { // 空格的个数
			fmt.Print(" ")
		}

		for j :=1 ; j <=2 * i - 1; j++ {  //*的个数
			if j == 1 || j == 2 * i - 1 {
				fmt.Printf("*")
			} else if i == 9 {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
			
		}
		fmt.Println("")

	}
}