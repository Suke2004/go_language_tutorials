package main

import "fmt"

func main() {

	//while loop
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i+=1
	}

	// //infinite loop
	// for{
	// 	fmt.Println("infinite")
	// }



	//classic for loop
	for j := 1; j < 5; j++ {
		fmt.Println(j)
	}

	///range
	for k := range 10 {
		fmt.Println(k) //runds from 0 to <10
	}

}
