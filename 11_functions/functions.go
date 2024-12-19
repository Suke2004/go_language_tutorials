package main

import (
	"fmt"
)
//func add (a,b int)int{ this is also fine if a,b are same data type
func add(a int,b int)int{//func name(arg,arg...)return data type{}
	return a+b
}

func getLang()(string,string,string){
	return "golang","javascript","C"
}

func processIt(fn func(a int)int){
	fn(2)
}

func processIT()func(a int)int{
	return func(a int)int{
		return 4
	}
}

func main(){
	fmt.Println(add(3,5))
	// lang1,lang2,lang3 := getLang()
	fmt.Println(getLang())

	fn := func(a int)int{
		return 2
	}
	println(processIt(fn))

	fn1:=processIT()
	fn1(6)
}