package basics

import (
	"fmt"
	"letsgo/arrays"
)

func Variables() {
	/* Variables in go*/
	fmt.Println("*******************" + "variables in go" + "*******************")
	var numbers int8 = 12
	var name string = "Deepak"
	fruit := "guava"
	fmt.Println("Hello world")
	fmt.Printf("%v\n%v\n%v\n", numbers, name, fruit)
	arrays.Arrays()
}
