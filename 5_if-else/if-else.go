package maim

import (
	"fmt"
)

func main() {
	age := 20
	if age >= 18 {
		fmt.Printf("Adult")
	} else if age >= 12 {
		fmt.Printf("teenager")
	} else {
		fmt.Printf("child")
	}

	if num := 11; num >= 10 {
		fmt.Printf("num is greater than 10")
	} else {
		fmt.Printf("Number is less than 10")
	}
}
