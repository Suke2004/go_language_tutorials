package main

import (
	"fmt"
	"time"
)


type customer struct{
	name string
	phone string
}
//struct,We can have methods to id,constructors to it,instances also
//by default and instance will have zero values if not assigned
type order struct{
	id string
	amount float32
	status string
	createdAt time.Time
	customer
}


func newOrder(id string,amount float32,status string) *order{
	myOrder := order {
		id: id,
		amount : amount,
		status: status,
		createdAt : time.Now(),
	}
	return &myOrder
}

func(o *order) changStatus(status string){
	o.status = status//you dont need to dereference as struct automatically dereferences and makes change in original instance.if you dont give pointer in function then it makes change in copy of that instance but not original instance
}

func(o *order) getAmount(){
	fmt.Println(o.amount)
}

func(c *customer) changPhone(phone string){
	c.phone = phone
}

func main(){
	//let us make a instance to the struct
	// myOrder := order{
	// 	id: "1",
	// 	amount: 50.5,
	// 	status: "recieved",
	// 	createdAt: time.Now(),
	// }
	// fmt.Println(myOrder)
	// // myOrder.status = "delivered"//we can change value of a struch like this or we can make a method to change also
	// myOrder.changStatus("delivered")
	// fmt.Println(myOrder.status)
	// myOrder.getAmount()

	// myOder2 := newOrder("2",100,"making")//using a construtor made by us
	// fmt.Println("order 2: ",myOder2)

	newCustomer := customer{
		name: "Harry Potter",
		phone: "alhomora",
	}
	newOrder := order{
		id: "1",
		amount: 100.50,
		status: "glorified",
		createdAt: time.Now(),
		customer: newCustomer,
	}
	fmt.Println(newOrder)
	fmt.Println(newOrder.customer)
	newCustomer.changPhone("spongify")
	fmt.Println(newCustomer)

}
