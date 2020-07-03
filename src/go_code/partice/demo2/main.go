package main



import (
	"fmt"
	"math/rand"
	"time"
)



// 数组的应用案例


func main()  {
	// var arr [26]byte
	// for i := 0; i < 26; i++ {
	// 	arr[i] = 'A' + byte(i)
	// }
	// fmt.Printf("%c",arr)

	// var arr2 = [5]int{5,4,8,1,7}
	// var maxV int = arr2[0]
	// var maxI int = 0
	// for i :=1; i <len(arr2); i++ {
	// 	if maxV < arr2[i] {
	// 		maxV = arr2[i]
	// 		maxI = i 
	// 	}
	// }
	// fmt.Println(maxV,maxI)

	// var arr2 = [5]float64{1.0, 4.0, 3.3, 9.7}
	// var sum float64 =0.0
	// for _,v := range arr2 {
	// 	sum += v
	// }
	// fmt.Println(sum,sum/float64(len(arr2)))
	
	var arr3 [5]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr3); i++ {
		arr3[i] = rand.Intn(100)
	}
	fmt.Println(arr3)
	for c := 0; c < len(arr3)/2; c++ {
		var tem int 
		tem =arr3[len(arr3)-c-1]
		arr3[len(arr3)-(c+1)] = arr3[c]
		arr3[c]=tem
	} 
	fmt.Println(arr3)
		
	

}