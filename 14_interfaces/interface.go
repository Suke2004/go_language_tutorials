package main

import(
	"fmt"
)

type paymenter interface{
	pay(amount float32)
}

type payment struct{
	// gaterazorpayway
	// gateway slice
	gateway paymenter
}

func (p *payment) makePayment(amount float32){
	// razorpayPaymentGw := razorpay{}
	// razorpayPaymentGw.pay(amount)
	// slicePaymentGw := slice{}
	// slicePaymentGw.pay(amount)
	//to overcome above changings we modified this
	p.gateway.pay(amount)
}

type razorpay struct{}

func(r *razorpay) pay(amount float32){
	//payment logic come here
	fmt.Println("payment using Razorpay",amount)
}

type slice struct{}

func(s slice) pay(amount float32){
	//payment logic come here
	fmt.Println("payment using Slice",amount)
}

type fake struct{}

func(f fake) pay(amount float32){
	//payment logic come here
	fmt.Println("payment using fakeGW",amount)
}

func main(){
	//still we need to change gateway here and in struct to make payment
	//this is open and cloce method which is not good for scalability
	// razorpayPaymentGw := razorpay{}
	// newPayment := payment{
	// 	gateway : razorpayPaymentGw,
	// }//This kind of changing always makes complex
	// slicePaymentGw := slice{}
	// newPayment := payment{
	// 	gateway: slicePaymentGw,
	// }
	fakeGw := slice{}
	newPayment := payment{
		gateway : fakeGw,//the methods with pointer will show error.To test that i kept pointer for razorpay,Test it
	}
	newPayment.makePayment(100)
}
