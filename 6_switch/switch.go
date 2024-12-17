package main

import(
	"fmt"
)

func main(){
	var i int = 7
	switch i{
	case 1:
		fmt.Printf("one")
	case 2:
		fmt.Printf("Two")
	default:
		fmt.Printf("Others")
	}
}