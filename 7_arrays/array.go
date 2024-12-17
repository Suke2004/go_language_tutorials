package main

import (
	"fmt"
)

func main(){
	//fixed length,static,constant time access

	var i [3]int
	i[0],i[1],i[2]=1,2,2
	j := [3]int{3,5,7}
	k := [3][3]int{{1,2,3},{4,5,6},{7,8,9}}
	var name [3]string
	name[0] = "string1"
	//int->0,string->"",boolean->false for unassigned indexes
	fmt.Println(i,j,k,name,len(j))

}
