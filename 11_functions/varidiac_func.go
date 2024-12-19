package main

import (
	"fmt"
)

// ...int means we can pass infinite number of arguments there of that datatype
func sum(nums ...int)int{
	sum := 0
	for _,num:=range nums{
		sum+=num
	}
	return sum
}

func main(){
	nums := []int{1,2,3,4,5,6}
	result := sum(nums...)//nums... is spread operator same like JS
	fmt.Printf("%d",result)
}