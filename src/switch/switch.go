package main

import "fmt"

func main() {
	var i int
	for {
		if i == 5 {
			i++
			continue
		}
		if i == 6 {
			break
		}
		fmt.Println(i)
		i++
	}
}
