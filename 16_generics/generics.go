package main

import(
	"fmt"
)

// func printSlice[T any] (items []T){
// 	for _,item := range items{
// 		fmt.Println(item)
// 	}
// }

// func printSlice[T interface{}] (items []T){//this and above fn allow any kind of variable which is not good always
// 	for _,item := range items{
// 		fmt.Println(item)
// 	}
// }

// func printSlice[T int | bool | string] (items []T){//this allows only int,bool,string
// 	for _,item := range items{
// 		fmt.Println(item)
// 	}
// }

// func printSlice[T comparable] (items []T){
// 	for _,item := range items{
// 		fmt.Println(item)
// 	}
// }

// type Stack[T any] struct{
// 	element []T
// }

type Stack[T any, V any] struct{
	name T
	dob []V
}

// func printStringSlice(items []string){
// 	for _,item := range items{
// 		fmt.Println(item)
// 	}
// }

func main(){
	// nums := []int{1,2,3}
	// names := []string{"Harry Potter", "Hermoine granger", "Ron Weasely"}
	// printSlice(names)
	// printStringSlice(names)//this for every kind of arguement type we are repeting same function.To reduce the code repetetion we use generics

	// myStack := Stack[int]{
	// 	element: []int{16,03,2004},
	// }
	myStack := Stack[string,int]{
		name : "Aditi",
		dob: []int{16,03,2004},
	}
	fmt.Println(myStack)
}