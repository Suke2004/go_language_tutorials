package main

import "fmt"

// func emailSender(emailChan chan string, done chan bool) {
// 	defer func() { done <- true }()
// 	for email := range emailChan {
// 		fmt.Println("Sending email to", email)///this will show deadlock
// 		time.Sleep(time.Second * 2)
// 	}
// }

func main() {

	chan1 := make(chan int)
	chan2 := make(chan string)
	go func() {
		chan1 <- 10
	}()
	go func() {
		chan2 <- "Aditi"
	}()
	for i := 0; i < 2; i++ {
		select {
		case chan1Var := <-chan1:
			fmt.Println("recieved data from channel 1", chan1Var)
		case chan2Var := <-chan2:
			fmt.Println("recieved data from channel 2", chan2Var)
		}
	}

	// var emailChan chan string
	// done := make(chan bool)
	// go emailSender(emailChan, done)
	// for i := 0; i < 100; i++ {  ///this will show deadlock
	// 	emailChan <- fmt.Sprintf("%d email", i)
	// }
	// fmt.Println("done sending")
	// <-done
}
