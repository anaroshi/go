package main

import "fmt"

var (
	arr  = [15]int{1, 9, 2, 11, 14, 4, 8, 3, 5, 7, 4, 6, 12, 2, 10}
	temp = [15]int{}
)

func main() {

	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		idx := arr[i]
		temp[idx]++
	}

	idx := 0
	for i := 0; i < len(temp); i++ {
		for j := 0; j < temp[i]; j++ {
			arr[idx] = i
			idx++
		}
	}
	fmt.Println(arr)
}
