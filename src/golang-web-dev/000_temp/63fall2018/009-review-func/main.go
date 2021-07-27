package main

import "fmt"

func main() {
	s := fmt.Sprintln("This is the entry point beginning of the programming")
	fmt.Println(s)

	foo()

	a := bar("James Bond")
	fmt.Println(a)

	fmt.Println(`Program "about"
	to
	exit`)
}

func foo() {
	fmt.Println("Foo is here")
}

func bar(x string) string {
	return fmt.Sprintln(x, "is here")
}

// We will pass ARGUMENTS in to a function that has been defined with PARAMETERS
// func receiver identifier(parameters) return(s) {code}