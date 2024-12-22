package main

import (
	"fmt"
	"time"
)

func task(id int){
	fmt.Println("doint task",id)
}

func main(){
	for i := 0;i<10;i++{
		go task(i)
	}

	time.Sleep(time.Second * 2)//this needed for the goroutine to finish
}