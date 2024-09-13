package main

import (
	"fmt"
	"strings"
)

func main() {
	val := [][]int{
		{2, 1, 5},
		{9, 3, 0},
	}

	val1 := []int{
		2, 1, 5,
	}

	fmt.Println("value:", val)

	fmt.Println("val1", val1[1:])

	n := 6

	// for i := 0; i < n; i++ {
	// 	for j := 0; j < n-i-1; j++ {
	// 		fmt.Print(" ")
	// 	}

	// 	for j := 0; j <= i; j++ {
	// 		fmt.Print("#")
	// 	}

	// 	fmt.Println()
	// }

	for i := 1; i <= n; i++ {
		fmt.Println(strings.Repeat(" ", n-i) + strings.Repeat("#", i))
	}

}
