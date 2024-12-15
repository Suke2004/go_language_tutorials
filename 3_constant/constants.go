package main

import (
	"fmt"
)
	//we cant define a variable outside a functiioin with short handle
func main() {
	const name = "Aditi"
	//name = "some" throws error as we cant reassign a constant

	const(
		port = 5000
		host = "localhost"
	)

	fmt.Println(name,port,host)


}