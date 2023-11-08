package arrays

import "fmt"

func TwoDArray() {
	matrix := [][]int8{{0, 1, 2, 3, 4}, {5, 6, 7, 8, 9}}

	for _, value := range matrix {
		for _, val := range value {
			fmt.Printf("%v ", val)
		}
		fmt.Printf("\n")
	}
}
