package main

import "fmt"

func main() {
	intSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, el := range intSlice {
		if el%2 == 0 {
			fmt.Println(el, "is even.")
		} else {
			fmt.Println(el, "is odd.")
		}
	}
}
