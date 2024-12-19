package main

import (
	"fmt"
)

func counter() func()int{
	count := 0

	return func()int{
		count+=1
		return count
	}
}

//closure means simply like C we dont have static variable here.To preserve the variable we use that kind of closure like we are closing the variable in a function
func main(){
	increment := counter()

	fmt.Println(increment())
	fmt.Println(increment())
}