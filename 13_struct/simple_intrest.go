package main

import(
	"fmt"
)


type person struct{
	name string
	phone string
	email string
}

type PTR struct{
	principle float32
	time float32//year
	rate float32//percentage
	person
}



func(p *PTR) getSimpleIntrest(){
	intrest := (p.principle * p.time * p.rate)/100
	fmt.Println(intrest)
}
func main(){

	newPerson := person{
		name : "akki",
		phone: "9999999999",
		email: "someemail@gmail.com",
	}
	newPTR := PTR{
		principle: 1000,
		time: 2,
		rate: 5,
		person: newPerson,
	}

	newPTR.getSimpleIntrest()

}
