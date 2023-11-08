package arrays

import (
	"fmt"
)

const SIZE int = 12

// var array_name = [length]datatype{values}
// var array_name = [...]datatype{values}

var fruits = [...]string{"guava", "apple", "pineapple", "mangoes", "papaya", "grapes", "dragonfruits"}

func Arrays() {
	// fmt.Println(fruits)
	for index, fruit := range fruits {
		fmt.Printf("%v fruit: %v\n", index+1, fruit)
	}
}
