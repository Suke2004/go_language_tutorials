package main

import (
	"fmt"
	"time"
)

func main(){
	i := time.Now().Weekday()

	switch i{
	case time.Saturday,time.Sunday:
		fmt.Printf("Weekend")
	default:
		fmt.Printf("Weekday")
	}
}
