package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int
	i = 16
	var j float32 //float64 also ok
	var k string = "Aditi"
	j = float32(i)
	fmt.Printf("%v %.2f %s %T %T %T\n", i, j, k, i, j, k)

	var (
		varA = 12
		varB = 34
	)
	vara := 12
	varb := 34
	fmt.Printf("%v %d %d %v\n", varA, varB, vara, varb)

	var n complex64 = 4+20i
	var n1 complex128 = 3+16i
	fmt.Printf("%v %v \n%T %T\n", n, n1,n,n1)

	//Bitwise
	a := 10
	b := 3
	fmt.Printf("%v\n", a & b)//AND
	fmt.Printf("%v\n", a | b)//OR
	fmt.Printf("%v\n", a ^ b)//XOR
	fmt.Printf("%v\n", a &^ b)//NOT
	fmt.Printf("%v\n", a << b)//a* 2^b
	fmt.Printf("%v\n", a >> b)// a/2^b

	//Strings
	s := "Hello "
	s2 := "World"

	fmt.Printf("%s\n",s+s2)///concatenation
	//Strings are not mutable
	//s[0] = 'A'//error
	byte := []byte(s)
	fmt.Printf("%v\n",byte)//prints ascii values

	var l string
	//l = string(i) throws error so...
	l = strconv.Itoa(i)
	fmt.Println(l)

}
