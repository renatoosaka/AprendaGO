package main

import (
	"fmt"
)

func main() {
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	for _, value := range slice {
		fmt.Println(value)
	}

	fmt.Printf("%T\n", slice)
}
