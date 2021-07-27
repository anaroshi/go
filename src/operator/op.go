package main

import (
	"fmt"
)

func main() {
	a := 1
	if a < 10 && a > 2 {
		fmt.Println("a is less than 10 and more than 2")
	} else {
		fmt.Println("a is more than 10 and less than 2")
	}
	fmt.Println("value of a is", a)

}
