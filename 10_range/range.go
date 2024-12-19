package main
import(
	"fmt"
)

func main(){
	//Iterating over data structures
	nums := []int{6,7,8}
	//range returns index,number
	for i,num := range nums{
		fmt.Println(i,num)
	}
	// for _,num := range nums{
	// 	fmt.Println(i,num)
	// }

	//returns key,value
	m := map[string]string{"fname":"Harry","lname":"Potter"}
	for k,v := range m{
		fmt.Println(k,v)
	}

	//unicode code point
	//starting byte of rune
	for i,c := range "golang"{
		fmt.Println(i,c)//returns the ascii value
		fmt.Println(i,string(c))//returns the string value corresponding to that ascii
	}



}
