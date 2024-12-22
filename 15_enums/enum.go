package main

import (
	"fmt"
)

// func changeOrderStatus(status string){
// 	fmt.Println("Changing order status to",status)
// }

// type Orderstatus int

// // const(
// // 	Recieved Orderstatus = iota
// // 	Confirmed
// // 	Prepared
// // 	Delivered
// // )

type Orderstatus string

const (
	Recieved  Orderstatus = "recieved"
	Confirmed             = "confirmed"
	Prepared              = "prepared"
	Delivered             = "delivered"
)

func changeOrderStatus(status Orderstatus) {
	fmt.Println("order status changing to ", status)
}

func main() {
	// changeOrderStatus("Recieved")//passing this many times may make typos sometimes in real world.To solve that enums came to play
	changeOrderStatus(Prepared) //this reduces typos and easy to export those to other files,and we can see the types of status there
}
