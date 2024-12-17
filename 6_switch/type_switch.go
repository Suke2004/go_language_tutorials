package main

import(
	"fmt"
)

func main(){

	i := "hello"
	// WhoAmI := func(i interface{}){
	// 	switch t := i.(type){
	// 	case int:
	// 		fmt.Printf("interger")
	// 	case string:
	// 		fmt.Printf("String")
	// 	default:
	// 		fmt.Printf("Other",t)
	// 	}
	// }

	WhoAmI := func(i interface{}){//we can use upper one also.anything is fine
		switch i.(type){
		case int:
			fmt.Printf("interger")
		case string:
			fmt.Printf("String")
		default:
			fmt.Printf("Other")
		}
	}
	WhoAmI(i)


}
