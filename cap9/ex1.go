package main

import (
	"fmt"
)

func main() {
	array := [5]int{0, 1, 2, 3, 4}

	for _, value := range array {
		fmt.Println(value)
	}

	fmt.Printf("%T\n", array)
}
