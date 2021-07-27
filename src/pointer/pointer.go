package main

import "fmt"

var (
	a    int
	memo = "hello"
)

func main() {

	a = 1

	fmt.Println(a)
	fmt.Println(memo)
	increase(&a)
	showMemo(&memo)

	fmt.Println(a)
	fmt.Println(memo)
}

func increase(x *int) {
	*x = *x + 1
}

func showMemo(str *string) {
	*str = "happy"
}
