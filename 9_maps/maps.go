package main

import (
	"fmt"
	"maps"
)

func main(){
	m:=make(map[string]string)
	m["name"] = "Aditi"
	m["area"] = "Pal"
	fmt.Println(m["name"],m["area"])
	//if key doesnt exist in map then it returns zero value to corresponding data type
	m["place"] = "Srik"

	delete(m,"area")//deletes the particular element
	fmt.Println(m)
	clear(m)//empty the map

	m1 := map[string]string{
		"name":"Aditi",
		"place":"Srik"}

	fmt.Println(maps.Equal(m,m1))


}

