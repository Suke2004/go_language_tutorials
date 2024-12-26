package main

import (
	"fmt"
	// "time"
)

// func processNum(numChan chan int){
// 	fmt.Println("processing number",<-numChan)
// }
// func processNum(numChan chan int){
// 	for num := range numChan{
// 		fmt.Println("processing number",num)
// 		time.Sleep(time.Second)
// 	}
// }

func sumresult (result chan int,num1 int,num2 int){
	numResult := num1+num2
	result <- numResult
}
func main(){
	result := make(chan int)
	go sumresult(result,4,5)
	res := <-result//blocking
	fmt.Println(res)
	

	// numChan := make(chan int)
	// go processNum(numChan)
	// for i := 0;i<100;i++{
	// 	numChan <- i
	// }
	// close(numChan)//this will close this numChan so that no blocking happens for other channel

	// go processNum(numChan)//here this process will run in other thread,channel will be assigned in other thread.This wont block
	// numChan <- 6//blocking


	// time.Sleep(time.Second * 2)


	// messageChan := make(chan string)// var messageChan chan string this is also ok
	// messageChan <- "ping"
	// msg := <-messageChan//blocking
	// fmt.Println(msg)
}