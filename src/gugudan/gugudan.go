package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {

		for j := 0; j < 4-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 0; i < 4; i++ {

		for j := 0; j < i+1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 7-i*2; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}

}
